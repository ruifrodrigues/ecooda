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

			query := NewQuery(s.conf)

			offset := api.Offset(req.GetPage(), req.GetPageSize(), maxLimit)
			limit := api.Limit(req.GetPageSize(), maxLimit)
			order := api.Sort(req.GetSort(), allowedSorts)

			records, err := query.GetLocations(offset, limit, order)
			if err != nil {
				return nil, err
			}

			count := query.CountLocations()
			cursor := api.NewCursor(count, req.GetPage(), req.GetPageSize(), maxLimit)

			collection := &pb.GetLocationCollectionResponse{}

			fields := api.RequestedFields(req.GetFields(), s.fields.Location)
			collection.Data = NewData(s.conf, records, fields, maxLimit).Generate()

			fields = s.fields.Location.Available
			collection.Meta = api.NewMeta().AddCursor(cursor).AddFields(fields).Meta

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

			record, err := NewQuery(s.conf).GetLocationByUuid(uuid, "*")
			if err != nil {
				return nil, err
			}

			requestedFields := api.RequestedFields(req.GetFields(), s.fields.Location)
			availableFields := s.fields.Location.Available

			item := new(pb.GetLocationItemResponse)
			item.Data = NewData(s.conf, []*Location{record}, requestedFields, 1).Generate()[0]
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

			query := NewQuery(s.conf)

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

			request := &pb.GetChallengeItemsBatchRequest{
				Uuids: challengeUuids,
				OptionalFields: &pb.GetChallengeItemsBatchRequest_Fields{
					Fields: req.GetFields(),
				},
				OptionalSort: &pb.GetChallengeItemsBatchRequest_Sort{
					Sort: req.GetSort(),
				},
				OptionalPageSize: &pb.GetChallengeItemsBatchRequest_PageSize{
					PageSize: req.GetPageSize(),
				},
				OptionalPage: &pb.GetChallengeItemsBatchRequest_Page{
					Page: req.GetPage(),
				},
			}

			batchResponse, err := NewGrpcClient(s.conf.Values.GrpcPort).GetChallengeItemsBatch(request)
			if err != nil {
				return nil, err
			}

			count := query.CountChallenges(location)
			cursor := api.NewCursor(count, req.GetPage(), req.GetPageSize(), maxLimit)
			fields := []string{batchResponse.GetMeta().GetFields()}

			_ = grpc.SetHeader(ctx, metadata.Pairs("x-http-code", "200"))

			return &pb.GetChallengesFromLocationResponse{
				Data: batchResponse.Data,
				Meta: api.NewMeta().AddCursor(cursor).AddFields(fields).Meta,
			}, nil
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

			query := NewQuery(s.conf)

			if req.GetParentUuid() != "" {
				record, err := query.GetLocationByUuid(req.GetParentUuid(), "id")
				if err != nil {
					return nil, err
				}
				fields["parent_id"] = record.ID
			}

			if err := query.CreateLocation(fields); err != nil {
				return nil, err
			}

			_ = grpc.SetHeader(ctx, metadata.Pairs("x-http-code", "201"))

			return &pb.CreateLocationResponse{}, nil
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

			return &pb.UpdateLocationResponse{}, nil
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

			if err := NewQuery(s.conf).DeleteLocation(uuid); err != nil {
				return nil, err
			}

			_ = grpc.SetHeader(ctx, metadata.Pairs("x-http-code", "200"))

			return &pb.DeleteLocationResponse{}, nil
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

			return &pb.AddCountryToLocationResponse{}, nil
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

			return &pb.AddRegionToLocationResponse{}, nil
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

			return &pb.RemoveCountryFromLocationResponse{}, nil
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

			return &pb.RemoveRegionFromLocationResponse{}, nil
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

			return &pb.AddChallengeToLocationResponse{}, nil
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

			return &pb.RemoveChallengeFromLocationResponse{}, nil
		}
	}
}
