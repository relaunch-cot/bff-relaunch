package handler

import (
	"context"

	"github.com/relaunch-cot/bff-relaunch/grpc"
	pb "github.com/relaunch-cot/lib-relaunch-cot/proto/post"
)

type IPost interface {
	CreatePost(ctx *context.Context, in *pb.CreatePostRequest) error
}

type postResource struct {
	grpcClient *grpc.Grpc
}

func (r *postResource) CreatePost(ctx *context.Context, in *pb.CreatePostRequest) error {
	err := r.grpcClient.PostGRPC.CreatePost(ctx, in)
	if err != nil {
		return err
	}

	return nil
}

func NewPostHandler(grpcClient *grpc.Grpc) IPost {
	return &postResource{
		grpcClient: grpcClient,
	}
}
