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
	newUser := &pb.User{
		Name:  in.Name,
		Email: in.Email,
	}

	return &pb.UpdateUserRequest{
		UserId:   in.UserId,
		Password: in.Password,
		NewUser:  newUser,
	}, nil
}

func UpdateUserPasswordToProto(email, currentPassword, newPassword string) (*pb.UpdateUserPasswordRequest, error) {
	return &pb.UpdateUserPasswordRequest{
		Email:           email,
		CurrentPassword: currentPassword,
		NewPassword:     newPassword,
	}, nil
}

func DeleteUserToProto(email, password string) (*pb.DeleteUserRequest, error) {
	return &pb.DeleteUserRequest{
		Email:    email,
		Password: password,
	}, nil
}
