package challenge

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

func (s *Service) GetChallengeItemsBatch(ctx context.Context, req *pb.GetChallengeItemsBatchRequest) (*pb.GetChallengeItemsBatchResponse, error) {
	ctx, cancel := context.WithTimeout(ctx, time.Millisecond*50)
	defer cancel()

	for {
		select {
		case <-ctx.Done():
			return nil, status.Error(codes.DeadlineExceeded, "Connection Timeout")
		default:
			var maxLimit int32 = 25
			var records []*Challenge

			uuids := req.GetUuids()
			for _, uuid := range uuids {
				if err := utils.ValidateUuid(uuid); err != nil {
					return nil, err
				}
			}

			offset := api.Offset(req.GetPage(), req.GetPageSize(), maxLimit)
			limit := api.Limit(req.GetPageSize(), maxLimit)
			order := api.Sort(req.GetSort(), allowedSorts)

			records, err := NewQuery(s.conf).GetChallenges(offset, limit, order, uuids)
			if err != nil {
				return nil, err
			}

			count := int32(len(req.GetUuids()))
			cursor := api.NewCursor(count, req.GetPage(), req.GetPageSize(), maxLimit)

			collection := &pb.GetChallengeItemsBatchResponse{}

			fields := api.RequestedFields(req.GetFields(), s.fields.Challenge)
			collection.Data = NewChallengeData(s.conf, records, fields, maxLimit).Generate()

			fields = s.fields.Challenge.Available
			collection.Meta = api.NewMeta().AddCursor(cursor).AddFields(fields).Meta

			_ = grpc.SetHeader(ctx, metadata.Pairs("x-http-code", "200"))

			return collection, nil
		}
	}
}
