package handler

import "github.com/relaunch-cot/bff/model"

type IUser interface {
	createUser() model.createUserPOST
}

type userResource struct {
	grpc *grpc.GRPC
}

func (r *userResource) createUser() models.createUserPOST {

}

func NewUserHandler(grpc *grpc.Grpc) IUser {
	return &userResource{
		grpc: grpc,
	}
}
