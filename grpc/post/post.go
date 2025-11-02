package post

import (
	"context"

	pbPost "github.com/relaunch-cot/lib-relaunch-cot/proto/post"
)

type resource struct {
	grpcClient pbPost.PostServiceClient
}

type IPostGRPC interface {
	CreatePost(ctx *context.Context, in *pbPost.CreatePostRequest) error
}

func (r *resource) CreatePost(ctx *context.Context, in *pbPost.CreatePostRequest) error {
	_, err := r.grpcClient.CreatePost(*ctx, in)
	if err != nil {
		return err
	}

	return nil
}

func NewPostGrpcClient(grpcClient pbPost.PostServiceClient) IPostGRPC {
	return &resource{
		grpcClient: grpcClient,
	}
}
