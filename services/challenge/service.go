package challenge

import (
	"context"
	"github.com/ruifrodrigues/ecooda/config"
	challengev1 "github.com/ruifrodrigues/ecooda/stubs/go/challenge/v1"
	cursorv1 "github.com/ruifrodrigues/ecooda/stubs/go/cursor/v1"
	metav1 "github.com/ruifrodrigues/ecooda/stubs/go/meta/v1"
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

func (s *Service) GetChallenges(ctx context.Context, req *challengev1.GetChallengesRequest) (*challengev1.GetChallengesResponse, error) {
	ctx, cancel := context.WithTimeout(ctx, time.Millisecond*50)
	defer cancel()

	response := func() (*challengev1.GetChallengesResponse, error) {
		conn, err := s.config.Database.OpenConnection()
		if err != nil {
			panic("Database Connection is not open!")
		}

		model := conn.Model(&Challenge{})

		// GET TOTAL COUNT ITEMS
		var count int64 = 0
		model.Count(&count)

		response := &challengev1.GetChallengesResponse{
			Data: transformer(model, req),
			Meta: &metav1.Meta{
				OptionalCursor: &metav1.Meta_Cursor{
					Cursor: &cursorv1.Cursor{
						Previous: 0,
						Current:  1,
						Next:     2,
						Count:    int32(count),
					},
				},
			},
		}

		return response, nil
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
