package location

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

func (s *Service) GetLocationCollection(ctx context.Context, req *pb.GetLocationCollectionRequest) (*pb.GetLocationCollectionResponse, error) {
	ctx, cancel := context.WithTimeout(ctx, time.Millisecond*50)
	defer cancel()

	for {
		select {
		case <-ctx.Done():
			return nil, status.Error(codes.DeadlineExceeded, "Connection Timeout")
		default:
			var maxLimit int32 = 25

			dbCtx := s.conf.Database.Ctx()

			query := NewQuery(dbCtx)
			count := query.CountLocations()
			cursor, err := api.NewCursor(count, req.GetPage(), req.GetPageSize(), maxLimit)
			if err != nil {
				return nil, err
			}

			offset := api.Offset(req.GetPage(), req.GetPageSize(), maxLimit)
			limit := api.Limit(req.GetPageSize(), maxLimit)
			order := api.Sort(req.GetSort(), allowedSorts)

			records, err := query.GetLocations(offset, limit, order)
			if err != nil {
				return nil, err
			}

			requestedFields := api.RequestedFields(req.GetFields(), s.fields.Location)
			availableFields := s.fields.Location.Available

			collection := new(pb.GetLocationCollectionResponse)
			collection.Data = locationFields(dbCtx, records, requestedFields, maxLimit)
			collection.Meta = api.NewMeta().AddCursor(cursor).AddFields(availableFields).Meta

			_ = grpc.SetHeader(ctx, metadata.Pairs("x-http-code", "200"))

			return collection, nil
		}
	}
}

func (s *Service) GetLocationItem(ctx context.Context, req *pb.GetLocationItemRequest) (*pb.GetLocationItemResponse, error) {
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

			dbCtx := s.conf.Database.Ctx()

			record, err := NewQuery(dbCtx).GetLocationByUuid(uuid, "*")
			if err != nil {
				return nil, err
			}

			requestedFields := api.RequestedFields(req.GetFields(), s.fields.Location)
			availableFields := s.fields.Location.Available

			item := new(pb.GetLocationItemResponse)
			item.Data = locationFields(dbCtx, []*Location{record}, requestedFields, 1)[0]
			item.Meta = api.NewMeta().AddFields(availableFields).Meta

			_ = grpc.SetHeader(ctx, metadata.Pairs("x-http-code", "200"))

			return item, nil
		}
	}
}

func (s *Service) GetChallengesFromLocation(ctx context.Context, req *pb.GetChallengesFromLocationRequest) (*pb.GetChallengesFromLocationResponse, error) {
	ctx, cancel := context.WithTimeout(ctx, time.Millisecond*50)
	defer cancel()

	for {
		select {
		case <-ctx.Done():
			return nil, status.Error(codes.DeadlineExceeded, "Connection Timeout")
		default:
			var maxLimit int32 = 25

			uuid := req.GetUuid()
			if err := utils.ValidateUuid(uuid); err != nil {
				return nil, err
			}

			query := NewQuery(s.conf.Database.Ctx())

			location, err := query.GetLocationByUuid(uuid, "*")
			if err != nil {
				return nil, err
			}

			offset := api.Offset(req.GetPage(), req.GetPageSize(), maxLimit)
			limit := api.Limit(req.GetPageSize(), maxLimit)
			order := api.Sort(req.GetSort(), allowedSorts)

			challengeUuids, err := query.GetChallengesUuids(location, offset, limit, order)
			if err != nil {
				return nil, err
			}

			request := new(pb.GetChallengeItemsBatchRequest)
			request.Uuids = challengeUuids

			optionalFields := new(pb.GetChallengeItemsBatchRequest_Fields)
			optionalFields.Fields = req.GetFields()
			request.OptionalFields = optionalFields

			optionalSort := new(pb.GetChallengeItemsBatchRequest_Sort)
			optionalSort.Sort = req.GetSort()
			request.OptionalSort = optionalSort

			optionalPageSize := new(pb.GetChallengeItemsBatchRequest_PageSize)
			optionalPageSize.PageSize = req.GetPageSize()
			request.OptionalPageSize = optionalPageSize

			optionalPage := new(pb.GetChallengeItemsBatchRequest_Page)
			optionalPage.Page = req.GetPage()
			request.OptionalPage = optionalPage

			batchResponse, err := NewGrpcClient(s.conf.Values.GrpcPort).GetChallengeItemsBatch(request)
			if err != nil {
				return nil, err
			}

			count := query.CountChallenges(location)
			cursor, err := api.NewCursor(count, req.GetPage(), req.GetPageSize(), maxLimit)
			if err != nil {
				return nil, err
			}

			availableFields := []string{batchResponse.GetMeta().GetFields()}

			collection := new(pb.GetChallengesFromLocationResponse)
			collection.Data = batchResponse.Data
			collection.Meta = api.NewMeta().
				AddCursor(cursor).
				AddFields(availableFields).
				Meta

			_ = grpc.SetHeader(ctx, metadata.Pairs("x-http-code", "200"))

			return collection, nil
		}
	}
}

func (s *Service) CreateLocation(ctx context.Context, req *pb.CreateLocationRequest) (*pb.CreateLocationResponse, error) {
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
				"type": req.GetType(),
			}

			dbCtx := s.conf.Database.Ctx()

			if req.GetParentUuid() != "" {
				record, err := NewQuery(dbCtx).GetLocationByUuid(req.GetParentUuid(), "id")
				if err != nil {
					return nil, err
				}
				fields["parent_id"] = record.ID
			}

			if err := NewQuery(dbCtx).CreateLocation(fields); err != nil {
				return nil, err
			}

			_ = grpc.SetHeader(ctx, metadata.Pairs("x-http-code", "201"))

			return new(pb.CreateLocationResponse), nil
		}
	}
}

func (s *Service) UpdateLocation(ctx context.Context, req *pb.UpdateLocationRequest) (*pb.UpdateLocationResponse, error) {
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
			aggregateRoot.ChangeType(req.GetType())

			if err = repository.Save(); err != nil {
				return nil, err
			}

			_ = grpc.SetHeader(ctx, metadata.Pairs("x-http-code", "200"))

			return new(pb.UpdateLocationResponse), nil
		}
	}
}

func (s *Service) DeleteLocation(ctx context.Context, req *pb.DeleteLocationRequest) (*pb.DeleteLocationResponse, error) {
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

			dbCtx := s.conf.Database.Ctx()

			if err := NewQuery(dbCtx).DeleteLocation(uuid); err != nil {
				return nil, err
			}

			_ = grpc.SetHeader(ctx, metadata.Pairs("x-http-code", "200"))

			return new(pb.DeleteLocationResponse), nil
		}
	}
}

func (s *Service) AddCountryToLocation(ctx context.Context, req *pb.AddCountryToLocationRequest) (*pb.AddCountryToLocationResponse, error) {
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

			locationUuid := req.GetLocationUuid()
			if err := utils.ValidateUuid(locationUuid); err != nil {
				return nil, err
			}

			args := NewCommandArgs().Add(locationUuid)
			cmd := NewAddCountryCommand(uuid, args)

			if err := NewCommandHandler(s.conf).Handle(cmd); err != nil {
				return nil, err
			}

			_ = grpc.SetHeader(ctx, metadata.Pairs("x-http-code", "200"))

			return new(pb.AddCountryToLocationResponse), nil
		}
	}
}

func (s *Service) AddRegionToLocation(ctx context.Context, req *pb.AddRegionToLocationRequest) (*pb.AddRegionToLocationResponse, error) {
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

			locationUuid := req.GetLocationUuid()
			if err := utils.ValidateUuid(locationUuid); err != nil {
				return nil, err
			}

			args := NewCommandArgs().Add(locationUuid)
			cmd := NewAddRegionCommand(uuid, args)

			if err := NewCommandHandler(s.conf).Handle(cmd); err != nil {
				return nil, err
			}

			_ = grpc.SetHeader(ctx, metadata.Pairs("x-http-code", "200"))

			return new(pb.AddRegionToLocationResponse), nil
		}
	}
}

func (s *Service) RemoveCountryFromLocation(ctx context.Context, req *pb.RemoveCountryFromLocationRequest) (*pb.RemoveCountryFromLocationResponse, error) {
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

			locationUuid := req.GetLocationUuid()
			if err := utils.ValidateUuid(locationUuid); err != nil {
				return nil, err
			}

			args := NewCommandArgs().Add(locationUuid)
			cmd := NewRemoveCountryCommand(uuid, args)

			if err := NewCommandHandler(s.conf).Handle(cmd); err != nil {
				return nil, err
			}

			_ = grpc.SetHeader(ctx, metadata.Pairs("x-http-code", "200"))

			return new(pb.RemoveCountryFromLocationResponse), nil
		}
	}
}

func (s *Service) RemoveRegionFromLocation(ctx context.Context, req *pb.RemoveRegionFromLocationRequest) (*pb.RemoveRegionFromLocationResponse, error) {
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

			locationUuid := req.GetLocationUuid()
			if err := utils.ValidateUuid(locationUuid); err != nil {
				return nil, err
			}

			args := NewCommandArgs().Add(locationUuid)
			cmd := NewRemoveRegionCommand(uuid, args)

			if err := NewCommandHandler(s.conf).Handle(cmd); err != nil {
				return nil, err
			}

			_ = grpc.SetHeader(ctx, metadata.Pairs("x-http-code", "200"))

			return new(pb.RemoveRegionFromLocationResponse), nil
		}
	}
}

func (s *Service) AddChallengeToLocation(ctx context.Context, req *pb.AddChallengeToLocationRequest) (*pb.AddChallengeToLocationResponse, error) {
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

			ChallengeUuid := req.GetChallengeUuid()
			if err := utils.ValidateUuid(ChallengeUuid); err != nil {
				return nil, err
			}

			args := NewCommandArgs().Add(ChallengeUuid)
			cmd := NewAddChallengeCommand(uuid, args)

			if err := NewCommandHandler(s.conf).Handle(cmd); err != nil {
				return nil, err
			}

			_ = grpc.SetHeader(ctx, metadata.Pairs("x-http-code", "200"))

			return new(pb.AddChallengeToLocationResponse), nil
		}
	}
}

func (s *Service) RemoveChallengeFromLocation(ctx context.Context, req *pb.RemoveChallengeFromLocationRequest) (*pb.RemoveChallengeFromLocationResponse, error) {
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

			ChallengeUuid := req.GetChallengeUuid()
			if err := utils.ValidateUuid(ChallengeUuid); err != nil {
				return nil, err
			}

			args := NewCommandArgs().Add(ChallengeUuid)
			cmd := NewRemoveChallengeCommand(uuid, args)

			if err := NewCommandHandler(s.conf).Handle(cmd); err != nil {
				return nil, err
			}

			_ = grpc.SetHeader(ctx, metadata.Pairs("x-http-code", "200"))

			return new(pb.RemoveChallengeFromLocationResponse), nil
		}
	}
}
