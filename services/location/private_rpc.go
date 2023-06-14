package location

import (
	"context"
	"github.com/ruifrodrigues/ecooda/api"
	pb "github.com/ruifrodrigues/ecooda/stubs/go/ecooda/v1"
	"github.com/ruifrodrigues/ecooda/utils"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"time"
)

func (s *Service) GetLocationFromChallenge(ctx context.Context, req *pb.GetLocationFromChallengeRequest) (*pb.GetLocationFromChallengeResponse, error) {
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

			item := new(pb.GetLocationFromChallengeResponse)

			record, err := NewQuery(s.conf).GetLocationsByChallengeUuid(uuid, "*")
			if err != nil {
				return item, err
			}

			requestedFields := api.RequestedFields(req.GetFields(), s.fields.Location)

			item.Data = NewData(s.conf, []*Location{record}, requestedFields, 1).Generate()[0]

			_ = grpc.SetHeader(ctx, metadata.Pairs("x-http-code", "200"))

			return item, nil
		}
	}
}
