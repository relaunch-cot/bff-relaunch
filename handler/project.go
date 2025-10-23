package handler

import (
	"context"

	"github.com/relaunch-cot/bff-relaunch/grpc"
	pb "github.com/relaunch-cot/lib-relaunch-cot/proto/project"
)

type IProject interface {
	CreateProject(ctx *context.Context, in *pb.CreateProjectRequest) error
	GetProject(ctx *context.Context, in *pb.GetProjectRequest) (*pb.GetProjectResponse, error)
	GetAllProjectsFromUser(ctx *context.Context, in *pb.GetAllProjectsFromUserRequest) (*pb.GetAllProjectsFromUserResponse, error)
	UpdateProject(ctx *context.Context, in *pb.UpdateProjectRequest) (*pb.UpdateProjectResponse, error)
	AddFreelancerToProject(ctx *context.Context, in *pb.AddFreelancerToProjectRequest) error
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

func (r *projectResource) GetAllProjectsFromUser(ctx *context.Context, in *pb.GetAllProjectsFromUserRequest) (*pb.GetAllProjectsFromUserResponse, error) {
	response, err := r.grpc.ProjectGRPC.GetAllProjectsFromUser(ctx, in)
	if err != nil {
		return nil, err
	}

	return response, err
}

func (r *projectResource) UpdateProject(ctx *context.Context, in *pb.UpdateProjectRequest) (*pb.UpdateProjectResponse, error) {
	response, err := r.grpc.ProjectGRPC.UpdateProject(ctx, in)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (r *projectResource) AddFreelancerToProject(ctx *context.Context, in *pb.AddFreelancerToProjectRequest) error {
	err := r.grpc.ProjectGRPC.AddFreelancerToProject(ctx, in)
	if err != nil {
		return err
	}

	return nil
}

func NewProjectHandler(grpc *grpc.Grpc) IProject {
	return &projectResource{
		grpc: grpc,
	}
}
