package challenge

import (
	"context"
	"github.com/google/uuid"
	"github.com/ruifrodrigues/ecooda/api"
	"github.com/ruifrodrigues/ecooda/utils"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"time"
)

func (s *Service) GetCategories(ctx Context, req *PbGetCategoriesRequest) (*PbGetCategoriesResponse, error) {
	ctx, cancel := context.WithTimeout(ctx, time.Millisecond*50)
	defer cancel()

	for {
		select {
		case <-ctx.Done():
			return nil, status.Error(codes.DeadlineExceeded, "Connection Timeout")
		default:
			var records []*Category

			request := api.NewCategoriesRequest(req)
			bag := api.NewRequestBag(request)
			offset := api.Offset(bag)
			limit := api.Limit(bag, 25)
			order := api.Sort(bag, allowedSorts)

			model := s.dbCtx.Model(&Category{})

			cursor, err := api.NewCursor(model, bag)
			if err != nil {
				return nil, err
			}

			s.dbCtx.Offset(offset).Limit(limit).Order(order).Find(&records)
			if len(records) == 0 {
				return nil, status.Error(codes.NotFound, "No records available")
			}

			requestedFields := api.RequestedFields(bag, s.fields.Category)
			availableFields := s.fields.Category.Available

			_ = grpc.SetHeader(ctx, metadata.Pairs("x-http-code", "200"))

			return &PbGetCategoriesResponse{
				Data: categoryFields(records, requestedFields, 25),
				Meta: api.NewMeta().AddCursor(cursor).AddFields(availableFields).Meta,
			}, nil
		}
	}
}

func (s *Service) GetCategory(ctx Context, req *PbGetCategoryRequest) (*PbGetCategoryResponse, error) {
	ctx, cancel := context.WithTimeout(ctx, time.Millisecond*25)
	defer cancel()

	for {
		select {
		case <-ctx.Done():
			return nil, status.Error(codes.DeadlineExceeded, "Connection Timeout")
		default:
			if !utils.ValidUUID(req.GetUuid()) {
				return nil, status.Error(codes.NotFound, "Invalid UUID")
			}

			var records []*Category

			result := s.dbCtx.Where("uuid=?", req.GetUuid()).Find(&records)
			if result.Error != nil && result.Error.Error() == "record not found" {
				return nil, status.Error(codes.NotFound, result.Error.Error())
			}

			if len(records) == 0 {
				return nil, status.Error(codes.NotFound, "Record does not exist")
			}

			request := api.NewCategoryRequest(req)
			bag := api.NewRequestBag(request)
			requestedFields := api.RequestedFields(bag, s.fields.Category)
			availableFields := s.fields.Category.Available

			_ = grpc.SetHeader(ctx, metadata.Pairs("x-http-code", "200"))

			return &PbGetCategoryResponse{
				Data: categoryFields(records, requestedFields, 1)[0],
				Meta: api.NewMeta().AddFields(availableFields).Meta,
			}, nil
		}
	}
}

func (s *Service) PostCategory(ctx Context, req *PbPostCategoryRequest) (*PbPostCategoryResponse, error) {
	ctx, cancel := context.WithTimeout(ctx, time.Millisecond*25)
	defer cancel()

	for {
		select {
		case <-ctx.Done():
			return nil, status.Error(codes.DeadlineExceeded, "Connection Timeout")
		default:
			if req.GetName() == "" {
				return nil, status.Error(codes.InvalidArgument, "Field Name is mandatory")
			}

			var fields = make(map[string]interface{})

			fields["uuid"] = uuid.New()
			fields["name"] = req.GetName()

			if len(fields) == 0 {
				return nil, status.Error(codes.InvalidArgument, "No fields provided")
			}

			result := s.dbCtx.Model(&Category{}).Create(fields)
			if result.Error != nil {
				return nil, status.Error(codes.AlreadyExists, result.Error.Error())
			}

			_ = grpc.SetHeader(ctx, metadata.Pairs("x-http-code", "201"))

			return &PbPostCategoryResponse{}, nil
		}
	}
}

func (s *Service) PutCategory(ctx Context, req *PbPutCategoryRequest) (*PbPutCategoryResponse, error) {
	ctx, cancel := context.WithTimeout(ctx, time.Millisecond*50)
	defer cancel()

	for {
		select {
		case <-ctx.Done():
			return nil, status.Error(codes.DeadlineExceeded, "Connection Timeout")
		default:
			if !utils.ValidUUID(req.GetUuid()) {
				return nil, status.Error(codes.NotFound, "Invalid UUID")
			}

			var fields = make(map[string]interface{})

			fields["name"] = req.GetName()

			if len(fields) == 0 {
				return nil, status.Error(codes.InvalidArgument, "No fields provided")
			}

			result := s.dbCtx.Model(&Category{}).Where("uuid=?", req.GetUuid()).Updates(fields)
			if result.Error != nil && result.Error.Error() == "record not found" {
				return nil, status.Error(codes.NotFound, result.Error.Error())
			}

			_ = grpc.SetHeader(ctx, metadata.Pairs("x-http-code", "204"))

			return &PbPutCategoryResponse{}, nil
		}
	}
}

func (s *Service) DeleteCategory(ctx Context, req *PbDeleteCategoryRequest) (*PbDeleteCategoryResponse, error) {
	ctx, cancel := context.WithTimeout(ctx, time.Millisecond*25)
	defer cancel()

	for {
		select {
		case <-ctx.Done():
			return nil, status.Error(codes.DeadlineExceeded, "Connection Timeout")
		default:
			if !utils.ValidUUID(req.GetUuid()) {
				return nil, status.Error(codes.NotFound, "Invalid UUID")
			}

			result := s.dbCtx.Where("uuid=?", req.GetUuid()).Delete(&Category{})
			if result.Error != nil && result.Error.Error() == "record not found" {
				return nil, status.Error(codes.NotFound, result.Error.Error())
			}

			_ = grpc.SetHeader(ctx, metadata.Pairs("x-http-code", "202"))

			return &PbDeleteCategoryResponse{}, nil
		}
	}
}
