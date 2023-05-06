package challenge

import (
	"context"
	"github.com/ruifrodrigues/ecooda/hateoas"
	pb "github.com/ruifrodrigues/ecooda/stubs/go/challenge/v1"
	"google.golang.org/genproto/googleapis/type/datetime"
)

type Service struct {
	pb.UnimplementedChallengeServiceServer
}

var api *hateoas.Hateoas

func init() {
	api = hateoas.New()
	api.AddCursor(cursor)
	api.AddLinks(links)
}

func NewChallengeService() *Service {
	return &Service{}
}

func (c *Service) GetChallenges(_ context.Context, _ *pb.GetChallengesRequest) (*pb.GetChallengesResponse, error) {
	challenges := []*pb.Challenge{
		{
			Uuid:        "62189ae8-e6cc-11ed-a05b-0242ac120003",
			Name:        "",
			Description: "description 1",
			Street:      "street 1",
			Postcode:    "2970-019",
			Location: &pb.Location{
				Continent: "Europe",
				Country:   "Portugal",
				Region:    "Lisbon Metro Area",
				City:      "Lisbon",
			},
			Latitude:  0.12131,
			Longitude: 1.203232,
			Thumbnail: "{}",
			Gallery:   "{}",
			CreatedAt: &datetime.DateTime{
				Year:    2023,
				Month:   5,
				Day:     5,
				Hours:   23,
				Minutes: 30,
				Seconds: 0,
			},
			UpdatedAt: &datetime.DateTime{
				Year:    2023,
				Month:   5,
				Day:     5,
				Hours:   23,
				Minutes: 30,
				Seconds: 0,
			},
		},
	}

	challenges = append(challenges, &pb.Challenge{})

	response := &pb.GetChallengesResponse{
		Data: challenges,
	}

	return response, nil
}
