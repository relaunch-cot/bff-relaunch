package transformer

import (
	models "github.com/relaunch-cot/bff-relaunch/models/user"
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
	baseUser := &pb.User{
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

func UpdateUserPasswordToProto(email, currentPassword, newPassword string) (*pb.UpdateUserPasswordRequest, error) {
	return &pb.UpdateUserPasswordRequest{
		Email:           email,
		CurrentPassword: currentPassword,
		NewPassword:     newPassword,
	}, nil
}
