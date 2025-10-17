package project

import (
	"context"

	pb "github.com/relaunch-cot/lib-relaunch-cot/proto/project"
)

type IProjectGRPC interface {
	CreateProject(ctx *context.Context, in *pb.CreateProjectRequest) error
}

type resource struct {
	grpcClient pb.ProjectServiceClient
}

func (r *resource) CreateProject(ctx *context.Context, in *pb.CreateProjectRequest) error {
	_, err := r.grpcClient.CreateProject(*ctx, in)
	if err != nil {
		return err
	}
	return nil
}

func NewProjectGrpcClient(grpcClient pb.ProjectServiceClient) IProjectGRPC {
	return &resource{
		grpcClient: grpcClient,
	}
}
