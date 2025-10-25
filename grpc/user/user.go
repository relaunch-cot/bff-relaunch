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
	DeleteUser(ctx *context.Context, in *pb.DeleteUserRequest) error
	GenerateReportFromJSON(ctx *context.Context, in *pb.GenerateReportRequest) (*pb.GenerateReportResponse, error)
	SendPasswordRecoveryEmail(ctx *context.Context, in *pb.SendPasswordRecoveryEmailRequest) error
	GetUserProfile(ctx *context.Context, in *pb.GetUserProfileRequest) (*pb.GetUserProfileResponse, error)
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

func (r *resource) DeleteUser(ctx *context.Context, in *pb.DeleteUserRequest) error {
	_, err := r.grpcClient.DeleteUser(*ctx, in)
	if err != nil {
		return err
	}

	return nil
}

func (r *resource) GenerateReportFromJSON(ctx *context.Context, in *pb.GenerateReportRequest) (*pb.GenerateReportResponse, error) {
	response, err := r.grpcClient.GenerateReportFromJSON(*ctx, in)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (r *resource) SendPasswordRecoveryEmail(ctx *context.Context, in *pb.SendPasswordRecoveryEmailRequest) error {
	_, err := r.grpcClient.SendPasswordRecoveryEmail(*ctx, in)
	if err != nil {
		return err
	}

	return nil
}

func (r *resource) GetUserProfile(ctx *context.Context, in *pb.GetUserProfileRequest) (*pb.GetUserProfileResponse, error) {
	response, err := r.grpcClient.GetUserProfile(*ctx, in)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func NewUserGrpcClient(grpcClient pb.UserServiceClient) IUserGRPC {
	return &resource{
		grpcClient: grpcClient,
	}
}
