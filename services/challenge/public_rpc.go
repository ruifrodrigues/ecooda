package challenge

import (
	"context"
	_uuid "github.com/google/uuid"
	"github.com/ruifrodrigues/ecooda/api"
	pb "github.com/ruifrodrigues/ecooda/stubs/go/ecooda/v1"
	"github.com/ruifrodrigues/ecooda/utils"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"time"
)

func (s *Service) GetChallengeCollection(ctx context.Context, req *pb.GetChallengeCollectionRequest) (*pb.GetChallengeCollectionResponse, error) {
	ctx, cancel := context.WithTimeout(ctx, time.Millisecond*50)
	defer cancel()

	for {
		select {
		case <-ctx.Done():
			return nil, status.Error(codes.DeadlineExceeded, "Connection Timeout")
		default:
			var maxLimit int32 = 25
			var records []*Challenge

			query := NewQuery(s.conf)
			offset := api.Offset(req.GetPage(), req.GetPageSize(), maxLimit)
			limit := api.Limit(req.GetPageSize(), maxLimit)
			order := api.Sort(req.GetSort(), allowedSorts)

			records, err := query.GetChallenges(offset, limit, order, nil)
			if err != nil {
				return nil, err
			}

			count := query.CountChallenges()
			cursor := api.NewCursor(count, req.GetPage(), req.GetPageSize(), maxLimit)

			collection := &pb.GetChallengeCollectionResponse{}

			fields := api.RequestedFields(req.GetFields(), s.fields.Challenge)
			collection.Data = NewChallengeData(s.conf, records, fields, maxLimit).Generate()

			fields = s.fields.Challenge.Available
			collection.Meta = api.NewMeta().AddCursor(cursor).AddFields(fields).Meta

			_ = grpc.SetHeader(ctx, metadata.Pairs("x-http-code", "200"))

			return collection, nil
		}
	}
}

func (s *Service) GetChallengeItem(ctx context.Context, req *pb.GetChallengeItemRequest) (*pb.GetChallengeItemResponse, error) {
	ctx, cancel := context.WithTimeout(ctx, time.Millisecond*25)
	defer cancel()

	for {
		select {
		case <-ctx.Done():
			return nil, status.Error(codes.DeadlineExceeded, "Connection Timeout")
		default:
			uuid := req.GetUuid()
			if err := utils.ValidateUuid(uuid); err != nil {
				return nil, err
			}

			record, err := NewQuery(s.conf).GetChallengeByUuid(uuid)
			if err != nil {
				return nil, err
			}

			item := &pb.GetChallengeItemResponse{}

			fields := api.RequestedFields(req.GetFields(), s.fields.Challenge)
			item.Data = NewChallengeData(s.conf, []*Challenge{record}, fields, 1).Generate()[0]

			fields = s.fields.Challenge.Available
			item.Meta = api.NewMeta().AddFields(fields).Meta

			_ = grpc.SetHeader(ctx, metadata.Pairs("x-http-code", "200"))

			return item, nil
		}
	}
}

func (s *Service) CreateChallenge(ctx context.Context, req *pb.CreateChallengeRequest) (*pb.CreateChallengeResponse, error) {
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

			var fields = map[string]interface{}{
				"uuid": _uuid.New(),
				"name": req.GetName(),
			}

			if req.GetDescription() != "" {
				fields["description"] = req.GetDescription()
			}

			if req.GetStreet() != "" {
				fields["street"] = req.GetStreet()
			}

			if req.GetPostcode() != "" {
				fields["postcode"] = req.GetPostcode()
			}

			if req.GetLatitude() != 0 {
				fields["latitude"] = req.GetLatitude()
			}

			if req.GetLongitude() != 0 {
				fields["longitude"] = req.GetLongitude()
			}

			if req.GetThumbnail() != "" {
				fields["thumbnail"] = req.GetThumbnail()
			}

			if req.GetGallery() != "" {
				fields["gallery"] = req.GetGallery()
			}

			if err := NewQuery(s.conf).CreateChallenge(fields); err != nil {
				return nil, err
			}

			_ = grpc.SetHeader(ctx, metadata.Pairs("x-http-code", "201"))

			return &pb.CreateChallengeResponse{}, nil
		}
	}
}

func (s *Service) UpdateChallenge(ctx context.Context, req *pb.UpdateChallengeRequest) (*pb.UpdateChallengeResponse, error) {
	ctx, cancel := context.WithTimeout(ctx, time.Millisecond*50)
	defer cancel()

	for {
		select {
		case <-ctx.Done():
			return nil, status.Error(codes.DeadlineExceeded, "Connection Timeout")
		default:
			uuid := req.GetUuid()
			if err := utils.ValidateUuid(uuid); err != nil {
				return nil, err
			}

			repository := NewRepository(s.conf)

			aggregateRoot, err := repository.Get(uuid)
			if err != nil {
				return nil, err
			}

			aggregateRoot.ChangeName(req.GetName())
			aggregateRoot.ChangeDescription(req.GetDescription())
			aggregateRoot.ChangeStreet(req.GetStreet())
			aggregateRoot.ChangePostcode(req.GetPostcode())
			aggregateRoot.ChangeLatitude(req.GetLatitude())
			aggregateRoot.ChangeLongitude(req.GetLongitude())
			aggregateRoot.ChangeThumbnail(req.GetThumbnail())

			if err := aggregateRoot.ChangeGallery(req.GetGallery()); err != nil {
				return nil, err
			}

			if err := repository.Save(); err != nil {
				return nil, err
			}

			_ = grpc.SetHeader(ctx, metadata.Pairs("x-http-code", "200"))

			return &pb.UpdateChallengeResponse{}, nil
		}
	}
}

func (s *Service) DeleteChallenge(ctx context.Context, req *pb.DeleteChallengeRequest) (*pb.DeleteChallengeResponse, error) {
	ctx, cancel := context.WithTimeout(ctx, time.Millisecond*25)
	defer cancel()

	for {
		select {
		case <-ctx.Done():
			return nil, status.Error(codes.DeadlineExceeded, "Connection Timeout")
		default:
			uuid := req.GetUuid()
			if err := utils.ValidateUuid(uuid); err != nil {
				return nil, err
			}

			if err := NewQuery(s.conf).DeleteChallenge(uuid); err != nil {
				return nil, err
			}

			_ = grpc.SetHeader(ctx, metadata.Pairs("x-http-code", "200"))

			return &pb.DeleteChallengeResponse{}, nil
		}
	}
}

func (s *Service) AddCategoryToChallenge(ctx context.Context, req *pb.AddCategoryToChallengeRequest) (*pb.AddCategoryToChallengeResponse, error) {
	ctx, cancel := context.WithTimeout(ctx, time.Millisecond*50)
	defer cancel()

	for {
		select {
		case <-ctx.Done():
			return nil, status.Error(codes.DeadlineExceeded, "Connection Timeout")
		default:
			uuid := req.GetUuid()
			if err := utils.ValidateUuid(uuid); err != nil {
				return nil, err
			}

			categoryUuid := req.GetCategoryUuid()
			if err := utils.ValidateUuid(categoryUuid); err != nil {
				return nil, err
			}

			args := NewCommandArgs().Add(categoryUuid)
			cmd := NewAddCategoryCommand(uuid, args)

			if err := NewCommandHandler(s.conf).Handle(cmd); err != nil {
				return nil, err
			}

			_ = grpc.SetHeader(ctx, metadata.Pairs("x-http-code", "200"))

			return &pb.AddCategoryToChallengeResponse{}, nil
		}
	}
}

func (s *Service) RemoveCategoryFromChallenge(ctx context.Context, req *pb.RemoveCategoryFromChallengeRequest) (*pb.RemoveCategoryFromChallengeResponse, error) {
	ctx, cancel := context.WithTimeout(ctx, time.Millisecond*50)
	defer cancel()

	for {
		select {
		case <-ctx.Done():
			return nil, status.Error(codes.DeadlineExceeded, "Connection Timeout")
		default:
			uuid := req.GetUuid()
			if err := utils.ValidateUuid(uuid); err != nil {
				return nil, err
			}

			categoryUuid := req.GetCategoryUuid()
			if err := utils.ValidateUuid(categoryUuid); err != nil {
				return nil, err
			}

			args := NewCommandArgs().Add(categoryUuid)
			cmd := NewRemoveCategoryCommand(uuid, args)

			if err := NewCommandHandler(s.conf).Handle(cmd); err != nil {
				return nil, err
			}

			_ = grpc.SetHeader(ctx, metadata.Pairs("x-http-code", "200"))

			return &pb.RemoveCategoryFromChallengeResponse{}, nil
		}
	}
}

func (s *Service) GetCategoryCollection(ctx context.Context, req *pb.GetCategoryCollectionRequest) (*pb.GetCategoryCollectionResponse, error) {
	ctx, cancel := context.WithTimeout(ctx, time.Millisecond*50)
	defer cancel()

	for {
		select {
		case <-ctx.Done():
			return nil, status.Error(codes.DeadlineExceeded, "Connection Timeout")
		default:
			var maxLimit int32 = 25
			var records []*Category

			query := NewQuery(s.conf)

			offset := api.Offset(req.GetPage(), req.GetPageSize(), maxLimit)
			limit := api.Limit(req.GetPageSize(), maxLimit)
			order := api.Sort(req.GetSort(), allowedSorts)

			records, err := query.GetCategories(offset, limit, order)
			if err != nil {
				return nil, err
			}

			collection := &pb.GetCategoryCollectionResponse{}

			fields := api.RequestedFields(req.GetFields(), s.fields.Category)
			collection.Data = NewCategoryData(s.conf, records, fields, maxLimit).Generate()

			count := query.CountCategories()
			cursor := api.NewCursor(count, req.GetPage(), req.GetPageSize(), maxLimit)

			fields = s.fields.Category.Available
			collection.Meta = api.NewMeta().AddCursor(cursor).AddFields(fields).Meta

			_ = grpc.SetHeader(ctx, metadata.Pairs("x-http-code", "200"))

			return collection, nil
		}
	}
}

func (s *Service) GetCategoryItem(ctx context.Context, req *pb.GetCategoryItemRequest) (*pb.GetCategoryItemResponse, error) {
	ctx, cancel := context.WithTimeout(ctx, time.Millisecond*25)
	defer cancel()

	for {
		select {
		case <-ctx.Done():
			return nil, status.Error(codes.DeadlineExceeded, "Connection Timeout")
		default:
			uuid := req.GetUuid()
			if err := utils.ValidateUuid(uuid); err != nil {
				return nil, err
			}

			record, err := NewQuery(s.conf).GetCategoryByUuid(uuid)
			if err != nil {
				return nil, err
			}

			item := &pb.GetCategoryItemResponse{}

			fields := api.RequestedFields(req.GetFields(), s.fields.Category)
			item.Data = NewCategoryData(s.conf, []*Category{record}, fields, 1).Generate()[0]

			fields = s.fields.Category.Available
			item.Meta = api.NewMeta().AddFields(fields).Meta

			_ = grpc.SetHeader(ctx, metadata.Pairs("x-http-code", "200"))

			return item, nil
		}
	}
}

func (s *Service) CreateCategory(ctx context.Context, req *pb.CreateCategoryRequest) (*pb.CreateCategoryResponse, error) {
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

			fields["uuid"] = _uuid.New()
			fields["name"] = req.GetName()

			if err := NewQuery(s.conf).CreateCategory(fields); err != nil {
				return nil, err
			}

			_ = grpc.SetHeader(ctx, metadata.Pairs("x-http-code", "201"))

			return &pb.CreateCategoryResponse{}, nil
		}
	}
}

func (s *Service) UpdateCategory(ctx context.Context, req *pb.UpdateCategoryRequest) (*pb.UpdateCategoryResponse, error) {
	ctx, cancel := context.WithTimeout(ctx, time.Millisecond*50)
	defer cancel()

	for {
		select {
		case <-ctx.Done():
			return nil, status.Error(codes.DeadlineExceeded, "Connection Timeout")
		default:
			uuid := req.GetUuid()
			if err := utils.ValidateUuid(uuid); err != nil {
				return nil, err
			}

			if req.GetName() == "" {
				return nil, status.Error(codes.InvalidArgument, "Field Name is mandatory")
			}

			var fields = map[string]interface{}{"name": req.GetName()}

			if err := NewQuery(s.conf).UpdateCategory(uuid, fields); err != nil {
				return nil, err
			}

			_ = grpc.SetHeader(ctx, metadata.Pairs("x-http-code", "200"))

			return &pb.UpdateCategoryResponse{}, nil
		}
	}
}

func (s *Service) DeleteCategory(ctx context.Context, req *pb.DeleteCategoryRequest) (*pb.DeleteCategoryResponse, error) {
	ctx, cancel := context.WithTimeout(ctx, time.Millisecond*25)
	defer cancel()

	for {
		select {
		case <-ctx.Done():
			return nil, status.Error(codes.DeadlineExceeded, "Connection Timeout")
		default:
			uuid := req.GetUuid()
			if err := utils.ValidateUuid(uuid); err != nil {
				return nil, err
			}

			if err := NewQuery(s.conf).DeleteCategory(uuid); err != nil {
				return nil, err
			}

			_ = grpc.SetHeader(ctx, metadata.Pairs("x-http-code", "200"))

			return &pb.DeleteCategoryResponse{}, nil
		}
	}
}
