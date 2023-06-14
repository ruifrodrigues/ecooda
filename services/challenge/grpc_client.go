package challenge

import (
	"context"
	pb "github.com/ruifrodrigues/ecooda/stubs/go/ecooda/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Client interface {
	GetLocationFromChallenge(request *pb.GetLocationFromChallengeRequest) *pb.GetLocationFromChallengeResponse
}

type GrpcClient struct {
	client pb.LocationServiceClient
	conn   *grpc.ClientConn
}

func NewGrpcClient(port string) Client {
	conn, err := grpc.Dial("localhost"+port, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic("Can not dial GRPC port")
	}

	return &GrpcClient{
		client: pb.NewLocationServiceClient(conn),
	}
}

func (grpc *GrpcClient) GetLocationFromChallenge(request *pb.GetLocationFromChallengeRequest) *pb.GetLocationFromChallengeResponse {
	item, err := grpc.client.GetLocationFromChallenge(context.Background(), request)
	if err != nil {
		return &pb.GetLocationFromChallengeResponse{}
	}

	return item
}
