package location

import (
	"context"
	pb "github.com/ruifrodrigues/ecooda/stubs/go/ecooda/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/status"
)

type Client interface {
	GetChallengeItem(request *pb.GetChallengeItemRequest) (*pb.GetChallengeItemResponse, error)
	GetChallengeItemsBatch(request *pb.GetChallengeItemsBatchRequest) (*pb.GetChallengeItemsBatchResponse, error)
}

type GrpcClient struct {
	client pb.ChallengeServiceClient
	conn   *grpc.ClientConn
}

func NewGrpcClient(port string) Client {
	conn, err := grpc.Dial("localhost"+port, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic("Can not dial GRPC port")
	}

	return &GrpcClient{
		client: pb.NewChallengeServiceClient(conn),
	}
}

func (grpc *GrpcClient) GetChallengeItem(request *pb.GetChallengeItemRequest) (*pb.GetChallengeItemResponse, error) {
	item, err := grpc.client.GetChallengeItem(context.Background(), request)
	if err != nil {
		return nil, status.Error(codes.NotFound, "Challenge not found >> "+err.Error())
	}

	return item, nil
}

func (grpc *GrpcClient) GetChallengeItemsBatch(request *pb.GetChallengeItemsBatchRequest) (*pb.GetChallengeItemsBatchResponse, error) {
	item, err := grpc.client.GetChallengeItemsBatch(context.Background(), request)
	if err != nil {
		return nil, status.Error(codes.NotFound, "Challenge not found >> "+err.Error())
	}

	return item, nil
}
