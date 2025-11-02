package post

import (
	"context"

	"github.com/golang/protobuf/ptypes/empty"
	pbPost "github.com/relaunch-cot/lib-relaunch-cot/proto/post"
)

type resource struct {
	grpcClient pbPost.PostServiceClient
}

type IPostGRPC interface {
	CreatePost(ctx *context.Context, in *pbPost.CreatePostRequest) error
	GetPost(ctx *context.Context, in *pbPost.GetPostRequest) (*pbPost.GetPostResponse, error)
	GetAllPosts(ctx *context.Context) (*pbPost.GetAllPostsResponse, error)
}

func (r *resource) CreatePost(ctx *context.Context, in *pbPost.CreatePostRequest) error {
	_, err := r.grpcClient.CreatePost(*ctx, in)
	if err != nil {
		return err
	}

	return nil
}

func (r *resource) GetPost(ctx *context.Context, in *pbPost.GetPostRequest) (*pbPost.GetPostResponse, error) {
	response, err := r.grpcClient.GetPost(*ctx, in)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (r *resource) GetAllPosts(ctx *context.Context) (*pbPost.GetAllPostsResponse, error) {
	response, err := r.grpcClient.GetAllPosts(*ctx, &empty.Empty{})
	if err != nil {
		return nil, err
	}

	return response, nil
}

func NewPostGrpcClient(grpcClient pbPost.PostServiceClient) IPostGRPC {
	return &resource{
		grpcClient: grpcClient,
	}
}
