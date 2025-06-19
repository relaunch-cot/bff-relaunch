package transformer

import (
	model "github.com/relaunch-cot/bff/model/user"
	pb "github.com/relaunch-cot/lib-relaunch-cot/proto/user"
)

func CreateUserToProto(in *model.User) (*pb.CreateUserRequest, error) {
	return &pb.CreateUserRequest{
		Name:     in.Name,
		Email:    in.Email,
		Password: in.Password,
	}, nil
}
