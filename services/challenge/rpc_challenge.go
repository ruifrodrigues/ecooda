package challenge

import (
	"context"
	"encoding/json"
	"github.com/google/uuid"
	"github.com/ruifrodrigues/ecooda/api"
	"github.com/ruifrodrigues/ecooda/utils"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"gorm.io/datatypes"
	"time"
)

func (s *Service) GetChallenges(ctx Context, req *PbGetChallengesRequest) (*PbGetChallengesResponse, error) {
	ctx, cancel := context.WithTimeout(ctx, time.Millisecond*50)
	defer cancel()

	for {
		select {
		case <-ctx.Done():
			return nil, status.Error(codes.DeadlineExceeded, "Connection Timeout")
		default:
			var maxLimit int32 = 25
			var records []*Challenge

			model := s.dbCtx.Model(&Challenge{})
			cursor, err := api.NewCursor(model, req.GetPage(), req.GetPageSize(), maxLimit)
			if err != nil {
				return nil, err
			}

			offset := api.Offset(req.GetPage(), req.GetPageSize(), maxLimit)
			limit := api.Limit(req.GetPageSize(), maxLimit)
			order := api.Sort(req.GetSort(), allowedSorts)

			s.dbCtx.Preload("Categories").Offset(offset).Limit(limit).Order(order).Find(&records)
			if len(records) == 0 {
				return nil, status.Error(codes.NotFound, "No records available")
			}

			requestedFields := api.RequestedFields(req.GetFields(), s.fields.Challenge)
			availableFields := s.fields.Challenge.Available

			_ = grpc.SetHeader(ctx, metadata.Pairs("x-http-code", "200"))

			return &PbGetChallengesResponse{
				Data: challengeFields(records, requestedFields, maxLimit),
				Meta: api.NewMeta().AddCursor(cursor).AddFields(availableFields).Meta,
			}, nil
		}
	}
}

func (s *Service) GetChallenge(ctx Context, req *PbGetChallengeRequest) (*PbGetChallengeResponse, error) {
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

			var records []*Challenge

			result := s.dbCtx.Preload("Categories").Where("uuid=?", req.GetUuid()).Find(&records)
			if result.Error != nil && result.Error.Error() == "record not found" {
				return nil, status.Error(codes.NotFound, result.Error.Error())
			}

			if len(records) == 0 {
				return nil, status.Error(codes.NotFound, "Record does not exist")
			}

			requestedFields := api.RequestedFields(req.GetFields(), s.fields.Challenge)
			availableFields := s.fields.Challenge.Available

			_ = grpc.SetHeader(ctx, metadata.Pairs("x-http-code", "200"))

			return &PbGetChallengeResponse{
				Data: challengeFields(records, requestedFields, 1)[0],
				Meta: api.NewMeta().AddFields(availableFields).Meta,
			}, nil
		}
	}
}

func (s *Service) PostChallenge(ctx Context, req *PbPostChallengeRequest) (*PbPostChallengeResponse, error) {
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

			result := s.dbCtx.Model(&Challenge{}).Create(fields)
			if result.Error != nil {
				return nil, status.Error(codes.AlreadyExists, result.Error.Error())
			}

			_ = grpc.SetHeader(ctx, metadata.Pairs("x-http-code", "201"))

			return &PbPostChallengeResponse{}, nil
		}
	}
}

func (s *Service) PutChallenge(ctx Context, req *PbPutChallengeRequest) (*PbPutChallengeResponse, error) {
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

			repository := NewRepository(s.dbCtx)
			aggregateRoot := repository.Get(req.GetUuid())

			if req.GetName() != "" {
				aggregateRoot.ChangeName(req.GetName())
			}

			if req.GetDescription() != "" {
				aggregateRoot.ChangeDescription(req.GetDescription())
			}

			if req.GetStreet() != "" {
				aggregateRoot.ChangeStreet(req.GetStreet())
			}

			if req.GetPostcode() != "" {
				aggregateRoot.ChangePostcode(req.GetPostcode())
			}

			if req.GetLatitude() != 0 {
				aggregateRoot.ChangeLatitude(req.GetLatitude())
			}

			if req.GetLongitude() != 0 {
				aggregateRoot.ChangeLongitude(req.GetLongitude())
			}

			if req.GetThumbnail() != "" {
				aggregateRoot.ChangeThumbnail(req.GetThumbnail())
			}

			if req.GetGallery() != "" {
				var gallery datatypes.JSON

				err := json.Unmarshal([]byte(req.GetGallery()), &gallery)
				if err != nil {
					return nil, status.Error(codes.InvalidArgument, "Gallery field not a valid JSON")
				}

				aggregateRoot.ChangeGallery(gallery)
			}

			if err := repository.Save(); err != nil {
				return nil, err
			}

			_ = grpc.SetHeader(ctx, metadata.Pairs("x-http-code", "204"))

			return &PbPutChallengeResponse{}, nil
		}
	}
}

func (s *Service) DeleteChallenge(ctx Context, req *PbDeleteChallengeRequest) (*PbDeleteChallengeResponse, error) {
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

			result := s.dbCtx.Where("uuid=?", req.GetUuid()).Delete(&Challenge{})
			if result.Error != nil && result.Error.Error() == "record not found" {
				return nil, status.Error(codes.NotFound, result.Error.Error())
			}

			_ = grpc.SetHeader(ctx, metadata.Pairs("x-http-code", "202"))

			return &PbDeleteChallengeResponse{}, nil
		}
	}
}

func (s *Service) CommandChallenge(ctx Context, req *PbCommandChallengeRequest) (*PbCommandChallengeResponse, error) {
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

			record := &Category{}

			result := s.dbCtx.Where("uuid=?", req.GetCategoryUuid()).First(record)
			if result.Error != nil && result.Error.Error() == "record not found" {
				return nil, status.Error(codes.NotFound, result.Error.Error())
			}

			repository := NewRepository(s.dbCtx)
			aggregateRoot := repository.Get(req.GetUuid())

			if req.GetCommand() == AddCategory {
				aggregateRoot.AddCategory(record)
				goto FINISH
			}

			if req.GetCommand() == RemoveCategory {
				aggregateRoot.RemoveCategory(req.GetCategoryUuid())
				goto FINISH
			}

		FINISH:
			if err := repository.Save(); err != nil {
				return nil, err
			}

			_ = grpc.SetHeader(ctx, metadata.Pairs("x-http-code", "202"))

			return &PbCommandChallengeResponse{}, nil
		}
	}
}
