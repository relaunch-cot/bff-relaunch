package project

import (
	"context"

	pb "github.com/relaunch-cot/lib-relaunch-cot/proto/project"
)

type IProjectGRPC interface {
	CreateProject(ctx *context.Context, in *pb.CreateProjectRequest) error
	GetProject(ctx *context.Context, in *pb.GetProjectRequest) (*pb.GetProjectResponse, error)
	GetAllProjectsFromUser(ctx *context.Context, in *pb.GetAllProjectsFromUserRequest) (*pb.GetAllProjectsFromUserResponse, error)
	UpdateProject(ctx *context.Context, in *pb.UpdateProjectRequest) (*pb.UpdateProjectResponse, error)
	AddFreelancerToProject(ctx *context.Context, in *pb.AddFreelancerToProjectRequest) error
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

func (r *resource) GetProject(ctx *context.Context, in *pb.GetProjectRequest) (*pb.GetProjectResponse, error) {
	resp, err := r.grpcClient.GetProject(*ctx, in)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (r *resource) GetAllProjectsFromUser(ctx *context.Context, in *pb.GetAllProjectsFromUserRequest) (*pb.GetAllProjectsFromUserResponse, error) {
	resp, err := r.grpcClient.GetAllProjectsFromUser(*ctx, in)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (r *resource) UpdateProject(ctx *context.Context, in *pb.UpdateProjectRequest) (*pb.UpdateProjectResponse, error) {
	resp, err := r.grpcClient.UpdateProject(*ctx, in)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (r *resource) AddFreelancerToProject(ctx *context.Context, in *pb.AddFreelancerToProjectRequest) error {
	_, err := r.grpcClient.AddFreelancerToProject(*ctx, in)
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
