package location

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"github.com/ruifrodrigues/ecooda/api"
	"github.com/ruifrodrigues/ecooda/utils"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"time"
)

func (s *Service) GetLocations(ctx Context, req *PbGetLocationsRequest) (*PbGetLocationsResponse, error) {
	ctx, cancel := context.WithTimeout(ctx, time.Millisecond*50)
	defer cancel()

	for {
		select {
		case <-ctx.Done():
			return nil, status.Error(codes.DeadlineExceeded, "Connection Timeout")
		default:
			var maxLimit int32 = 25
			var records []*Location

			model := s.dbCtx.Find(&records)
			cursor, err := api.NewCursor(model, req.GetPage(), req.GetPageSize(), maxLimit)
			if err != nil {
				return nil, err
			}

			offset := api.Offset(req.GetPage(), req.GetPageSize(), maxLimit)
			limit := api.Limit(req.GetPageSize(), maxLimit)
			order := api.Sort(req.GetSort(), allowedSorts)

			s.dbCtx.Offset(offset).Limit(limit).Order(order).Find(&records)
			if len(records) == 0 {
				return nil, status.Error(codes.NotFound, "No records available")
			}

			requestedFields := api.RequestedFields(req.GetFields(), s.fields.Location)
			availableFields := s.fields.Location.Available

			_ = grpc.SetHeader(ctx, metadata.Pairs("x-http-code", "200"))

			return &PbGetLocationsResponse{
				Data: locationFields(s.dbCtx, records, requestedFields, maxLimit),
				Meta: api.NewMeta().AddCursor(cursor).AddFields(availableFields).Meta,
			}, nil
		}
	}
}

func (s *Service) GetLocation(ctx Context, req *PbGetLocationRequest) (*PbGetLocationResponse, error) {
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

			var records []*Location

			result := s.dbCtx.Where("uuid=?", req.GetUuid()).Find(&records)
			if result.Error != nil && result.Error.Error() == "record not found" {
				return nil, status.Error(codes.NotFound, result.Error.Error())
			}

			if len(records) == 0 {
				return nil, status.Error(codes.NotFound, "Record does not exist")
			}

			requestedFields := api.RequestedFields(req.GetFields(), s.fields.Location)
			availableFields := s.fields.Location.Available

			_ = grpc.SetHeader(ctx, metadata.Pairs("x-http-code", "200"))

			return &PbGetLocationResponse{
				Data: locationFields(s.dbCtx, records, requestedFields, 1)[0],
				Meta: api.NewMeta().AddFields(availableFields).Meta,
			}, nil
		}
	}
}

func (s *Service) PostLocation(ctx Context, req *PbPostLocationRequest) (*PbPostLocationResponse, error) {
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
			fields["type"] = req.GetType()

			if req.GetParentUuid() != "" {
				var record *Location
				result := s.dbCtx.Select("id").Where("uuid=?", req.GetParentUuid()).Find(&record)
				if result.Error != nil && result.Error.Error() == "record not found" {
					return nil, status.Error(codes.NotFound, result.Error.Error())
				}

				fmt.Println(record)

				fields["parent_id"] = record.ID
			}

			result := s.dbCtx.Model(&Location{}).Create(fields)
			if result.Error != nil {
				return nil, status.Error(codes.AlreadyExists, result.Error.Error())
			}

			_ = grpc.SetHeader(ctx, metadata.Pairs("x-http-code", "201"))

			return &PbPostLocationResponse{}, nil
		}
	}
}
