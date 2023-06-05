package utils

import (
	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func ValidateUuid(u string) error {
	_, err := uuid.Parse(u)
	if err != nil {
		return status.Error(codes.InvalidArgument, "Invalid uuid")
	}

	return nil
}
