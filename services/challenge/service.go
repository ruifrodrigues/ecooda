package challenge

import (
	"context"
	"github.com/ruifrodrigues/ecooda/config"
	challengev1 "github.com/ruifrodrigues/ecooda/stubs/go/challenge/v1"
	cursorv1 "github.com/ruifrodrigues/ecooda/stubs/go/cursor/v1"
	locationv1 "github.com/ruifrodrigues/ecooda/stubs/go/location/v1"
	metav1 "github.com/ruifrodrigues/ecooda/stubs/go/meta/v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"time"
)

type (
	Context               = context.Context
	GetChallengesRequest  = challengev1.GetChallengesRequest
	GetChallengesResponse = challengev1.GetChallengesResponse

	PostChallengeRequest  = challengev1.PostChallengeRequest
	PostChallengeResponse = challengev1.PostChallengeResponse
)

type Service struct {
	challengev1.UnimplementedChallengeServiceServer
}

func NewChallengeService() *Service {
	return &Service{}
}

func (s *Service) GetChallenges(ctx Context, _ *GetChallengesRequest) (*GetChallengesResponse, error) {
	ctx, cancel := context.WithTimeout(ctx, time.Millisecond*50)
	defer cancel()

	uuid := "62189ae8-e6cc-11ed-a05b-0242ac120003"

	items := func() (*challengev1.GetChallengesResponse, error) {
		includes := []*challengev1.Include{
			{
				OptionalLocation: &challengev1.Include_Location{
					Location: &locationv1.Location{
						Uuid:      uuid,
						Continent: "Europe",
						Country:   "Portugal",
						Region:    "Lisbon Metro Area",
						City:      "Lisbon",
					},
				},
			},
		}

		challenges := []*challengev1.Challenge{
			{
				Uuid:        uuid,
				Name:        "Museu do Coches",
				Description: "description 1",
				Street:      "street 1",
				Postcode:    "2970-019",
				Latitude:    0.12131,
				Longitude:   1.203232,
				Thumbnail:   "{}",
				Gallery:     "{}",
				CreatedAt:   "2023-05-06 19:34:00",
				UpdatedAt:   "2023-05-06 19:34:00",
				Links:       config.NewLink(uuid, LinkList).List(),
				Includes:    includes,
			},
		}

		response := &GetChallengesResponse{
			Data: challenges,
			Meta: &metav1.Meta{
				OptionalCursor: &metav1.Meta_Cursor{
					Cursor: &cursorv1.Cursor{
						Previous: 0,
						Current:  1,
						Next:     2,
						Count:    40,
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
			return items()
		}
	}
}
