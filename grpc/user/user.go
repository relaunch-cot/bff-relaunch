package user

import (
	"context"
	pb "github.com/relaunch-cot/lib-relaunch-cot/proto/user"
)

type IUserGRPC interface {
	CreateUser(ctx *context.Context, user *pb.CreateUserRequest) error
}

type resource struct {
	grpcClient pb.UserServiceClient
}

func (r *resource) CreateUser(ctx *context.Context, in *pb.CreateUserRequest) error {
	_, err := r.grpcClient.CreateUser(*ctx, in)
	if err != nil {
		return err
	}

	return nil
}

func NewUserGrpcClient(grpcClient pb.UserServiceClient) IUserGRPC {
	return &resource{
		grpcClient: grpcClient,
	}
}
