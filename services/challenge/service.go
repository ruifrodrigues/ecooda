package challenge

import (
	"context"
	"github.com/ruifrodrigues/ecooda/config"
	"github.com/ruifrodrigues/ecooda/hateoas"
	challengev1 "github.com/ruifrodrigues/ecooda/stubs/go/challenge/v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"time"
)

type Service struct {
	challengev1.UnimplementedChallengeServiceServer
	config config.Config
}

func NewChallengeService(conf config.Config) *Service {
	return &Service{
		config: conf,
	}
}

func (s *Service) GetChallenges(
	ctx context.Context,
	req *challengev1.GetChallengesRequest,
) (
	*challengev1.GetChallengesResponse, error,
) {
	ctx, cancel := context.WithTimeout(ctx, time.Millisecond*50)
	defer cancel()

	response := func() (*challengev1.GetChallengesResponse, error) {
		conn, err := s.config.Database.OpenConnection()
		if err != nil {
			panic("Database Connection is not open!")
		}

		model := conn.Model(&Challenge{})

		request := hateoas.NewChallengesRequest(req)
		bag := hateoas.NewRequestBag(request)
		cursor := hateoas.NewCursor[*hateoas.ChallengeCollectionRequest](model, bag)
		meta := hateoas.NewMeta().AddCursor(cursor).Meta
		fields := hateoas.NewFields(DefaultFields, GuardedFields)
		transformer := hateoas.NewTransformer(model, bag, fields, MaxLimit)
		collection := hateoas.Collection(transformer, Fields, Includes)

		var data []*challengev1.Challenge
		for _, item := range collection {
			data = append(data, item.Challenge)
		}

		return &challengev1.GetChallengesResponse{
			Data: data,
			Meta: meta,
		}, nil
	}

	for {
		select {
		case <-ctx.Done():
			msg := "Connection Timeout"
			err := status.Error(codes.DeadlineExceeded, msg)

			return nil, err
		default:
			return response()
		}
	}
}

func (s *Service) GetChallenge(
	ctx context.Context,
	req *challengev1.GetChallengeRequest,
) (
	*challengev1.GetChallengeResponse, error,
) {
	ctx, cancel := context.WithTimeout(ctx, time.Millisecond*25)
	defer cancel()

	response := func() (*challengev1.GetChallengeResponse, error) {
		conn, err := s.config.Database.OpenConnection()
		if err != nil {
			panic("Database Connection is not open!")
		}

		model := conn.Model(&Challenge{})

		request := hateoas.NewChallengeRequest(req)
		bag := hateoas.NewRequestBag(request)
		fields := hateoas.NewFields(DefaultFields, GuardedFields)
		transformer := hateoas.NewTransformer(model, bag, fields, MaxLimit)
		item := hateoas.Item(transformer, Fields, Includes)

		return &challengev1.GetChallengeResponse{
			Data: item.Challenge,
		}, nil
	}

	for {
		select {
		case <-ctx.Done():
			msg := "Connection Timeout"
			err := status.Error(codes.DeadlineExceeded, msg)

			return nil, err
		default:
			return response()
		}
	}
}
