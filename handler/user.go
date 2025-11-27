package handler

import (
	"context"

	"github.com/relaunch-cot/bff-relaunch/grpc"
	pb "github.com/relaunch-cot/lib-relaunch-cot/proto/user"
)

type IUser interface {
	CreateUser(ctx *context.Context, in *pb.CreateUserRequest) error
	LoginUser(ctx *context.Context, in *pb.LoginUserRequest) (*pb.LoginUserResponse, error)
	UpdateUser(ctx *context.Context, in *pb.UpdateUserRequest) error
	UpdateUserPassword(ctx *context.Context, in *pb.UpdateUserPasswordRequest) error
	DeleteUser(ctx *context.Context, in *pb.DeleteUserRequest) error
	GenerateReportPDF(ctx *context.Context, in *pb.GenerateReportRequest) (*pb.GenerateReportResponse, error)
	SendPasswordRecoveryEmail(ctx *context.Context, in *pb.SendPasswordRecoveryEmailRequest) error
	GetUserProfile(ctx *context.Context, in *pb.GetUserProfileRequest) (*pb.GetUserProfileResponse, error)
	GetUserByName(ctx *context.Context, in *pb.GetUserByNameRequest) (*pb.GetUserByNameResponse, error)
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

func (r *userResource) UpdateUser(ctx *context.Context, in *pb.UpdateUserRequest) error {
	err := r.grpc.UserGRPC.UpdateUser(ctx, in)
	if err != nil {
		return err
	}

	return nil
}

func (r *userResource) UpdateUserPassword(ctx *context.Context, in *pb.UpdateUserPasswordRequest) error {
	err := r.grpc.UserGRPC.UpdateUserPassword(ctx, in)
	if err != nil {
		return err
	}

	return nil
}

func (r *userResource) DeleteUser(ctx *context.Context, in *pb.DeleteUserRequest) error {
	err := r.grpc.UserGRPC.DeleteUser(ctx, in)
	if err != nil {
		return err
	}

	return nil
}

func (r *userResource) GenerateReportPDF(ctx *context.Context, in *pb.GenerateReportRequest) (*pb.GenerateReportResponse, error) {
	response, err := r.grpc.UserGRPC.GenerateReportFromJSON(ctx, in)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (r *userResource) SendPasswordRecoveryEmail(ctx *context.Context, in *pb.SendPasswordRecoveryEmailRequest) error {
	err := r.grpc.UserGRPC.SendPasswordRecoveryEmail(ctx, in)
	if err != nil {
		return err
	}

	return nil
}

func (r *userResource) GetUserProfile(ctx *context.Context, in *pb.GetUserProfileRequest) (*pb.GetUserProfileResponse, error) {
	response, err := r.grpc.UserGRPC.GetUserProfile(ctx, in)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (r *userResource) GetUserByName(ctx *context.Context, in *pb.GetUserByNameRequest) (*pb.GetUserByNameResponse, error) {
	response, err := r.grpc.UserGRPC.GetUserByName(ctx, in)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func NewUserHandler(grpc *grpc.Grpc) IUser {
	return &userResource{
		grpc: grpc,
	}
}
