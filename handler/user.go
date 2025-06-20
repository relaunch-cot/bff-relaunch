package handler

import (
	"context"
	"github.com/relaunch-cot/bff/grpc"
	pb "github.com/relaunch-cot/lib-relaunch-cot/proto/user"
)

type IUser interface {
	CreateUser(ctx *context.Context, in *pb.CreateUserRequest) error
	LoginUser(ctx *context.Context, in *pb.LoginUserRequest) (*pb.LoginUserResponse, error)
}

type userResource struct {
	grpc *grpc.Grpc
}

func (r *userResource) CreateUser(ctx *context.Context, in *pb.CreateUserRequest) error {
	err := r.grpc.UserGRPC.CreateUser(ctx, in)
	if err != nil {
		return err
	}

	return nil
}

func (r *userResource) LoginUser(ctx *context.Context, in *pb.LoginUserRequest) (*pb.LoginUserResponse, error) {
	loginUserResponse, err := r.grpc.UserGRPC.LoginUser(ctx, in)
	if err != nil {
		return nil, err
	}

	return loginUserResponse, nil
}

func NewUserHandler(grpc *grpc.Grpc) IUser {
	return &userResource{
		grpc: grpc,
	}
}
