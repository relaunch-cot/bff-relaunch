package user

import (
	"context"

	pb "github.com/relaunch-cot/lib-relaunch-cot/proto/user"
)

type IUserGRPC interface {
	CreateUser(ctx *context.Context, user *pb.CreateUserRequest) error
	LoginUser(ctx *context.Context, in *pb.LoginUserRequest) (*pb.LoginUserResponse, error)
	UpdateUser(ctx *context.Context, in *pb.UpdateUserRequest) error
	UpdateUserPassword(ctx *context.Context, in *pb.UpdateUserPasswordRequest) error
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

func (r *resource) LoginUser(ctx *context.Context, in *pb.LoginUserRequest) (*pb.LoginUserResponse, error) {
	loginUserResponse, err := r.grpcClient.LoginUser(*ctx, in)
	if err != nil {
		return nil, err
	}

	return loginUserResponse, nil
}

func (r *resource) UpdateUser(ctx *context.Context, in *pb.UpdateUserRequest) error {
	_, err := r.grpcClient.UpdateUser(*ctx, in)
	if err != nil {
		return err
	}

	return nil
}

func (r *resource) UpdateUserPassword(ctx *context.Context, in *pb.UpdateUserPasswordRequest) error {
	_, err := r.grpcClient.UpdateUserPassword(*ctx, in)
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
