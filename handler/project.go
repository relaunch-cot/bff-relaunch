package handler

import (
	"context"

	"github.com/relaunch-cot/bff-relaunch/grpc"
	pb "github.com/relaunch-cot/lib-relaunch-cot/proto/project"
)

type IProject interface {
	CreateProject(ctx *context.Context, in *pb.CreateProjectRequest) error
	GetProject(ctx *context.Context, in *pb.GetProjectRequest) (*pb.GetProjectResponse, error)
}

type projectResource struct {
	grpc *grpc.Grpc
}

func (r *projectResource) CreateProject(ctx *context.Context, in *pb.CreateProjectRequest) error {
	err := r.grpc.ProjectGRPC.CreateProject(ctx, in)
	if err != nil {
		return err
	}

	return nil
}

func (r *projectResource) GetProject(ctx *context.Context, in *pb.GetProjectRequest) (*pb.GetProjectResponse, error) {
	response, err := r.grpc.ProjectGRPC.GetProject(ctx, in)
	if err != nil {
		return nil, err
	}

	return response, err
}

func NewProjectHandler(grpc *grpc.Grpc) IProject {
	return &projectResource{
		grpc: grpc,
	}
}
