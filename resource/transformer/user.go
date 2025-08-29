package transformer

import (
	models "github.com/relaunch-cot/bff-relaunch/models/user"
	pbBaseModels "github.com/relaunch-cot/lib-relaunch-cot/proto/base_models"
	pb "github.com/relaunch-cot/lib-relaunch-cot/proto/user"
)

func CreateUserToProto(in *models.User) (*pb.CreateUserRequest, error) {
	return &pb.CreateUserRequest{
		Name:     in.Name,
		Email:    in.Email,
		Password: in.Password,
	}, nil
}

func LoginUserToProto(in *models.User) (*pb.LoginUserRequest, error) {
	return &pb.LoginUserRequest{
		Email:    in.Email,
		Password: in.Password,
	}, nil
}

func UpdateUserToProto(in *models.User) (*pb.UpdateUserRequest, error) {
	baseUser := &pbBaseModels.User{
		UserId: in.UserId,
		Name:   in.Name,
		Email:  in.Email,
	}

	return &pb.UpdateUserRequest{
		Email:       in.Email,
		CurrentUser: baseUser,
		NewUser:     baseUser,
	}, nil
}

func UpdateUserDataToProto(in *models.User) (*pb.UpdateUserRequest, error) {
	baseUser := &pbBaseModels.User{
		UserId: in.UserId,
		Name:   in.Name,
		Email:  in.Email,
	}

	return &pb.UpdateUserRequest{
		Email:       in.Email,
		CurrentUser: baseUser,
		NewUser:     baseUser,
	}, nil
}
