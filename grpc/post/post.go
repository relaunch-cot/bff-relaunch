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
	GetAllPostsFromUser(ctx *context.Context, in *pbPost.GetAllPostsFromUserRequest) (*pbPost.GetAllPostsFromUserResponse, error)
	UpdatePost(ctx *context.Context, in *pbPost.UpdatePostRequest) (*pbPost.UpdatePostResponse, error)
	DeletePost(ctx *context.Context, in *pbPost.DeletePostRequest) error
	UpdateLikesFromPost(ctx *context.Context, in *pbPost.UpdateLikesFromPostRequest) (*pbPost.UpdateLikesFromPostResponse, error)
	AddCommentToPost(ctx *context.Context, in *pbPost.AddCommentToPostRequest) (*pbPost.AddCommentToPostResponse, error)
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

func (r *resource) GetAllPostsFromUser(ctx *context.Context, in *pbPost.GetAllPostsFromUserRequest) (*pbPost.GetAllPostsFromUserResponse, error) {
	response, err := r.grpcClient.GetAllPostsFromUser(*ctx, in)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (r *resource) UpdatePost(ctx *context.Context, in *pbPost.UpdatePostRequest) (*pbPost.UpdatePostResponse, error) {
	response, err := r.grpcClient.UpdatePost(*ctx, in)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (r *resource) DeletePost(ctx *context.Context, in *pbPost.DeletePostRequest) error {
	_, err := r.grpcClient.DeletePost(*ctx, in)
	if err != nil {
		return err
	}

	return nil
}

func (r *resource) UpdateLikesFromPost(ctx *context.Context, in *pbPost.UpdateLikesFromPostRequest) (*pbPost.UpdateLikesFromPostResponse, error) {
	response, err := r.grpcClient.UpdateLikesFromPost(*ctx, in)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (r *resource) AddCommentToPost(ctx *context.Context, in *pbPost.AddCommentToPostRequest) (*pbPost.AddCommentToPostResponse, error) {
	response, err := r.grpcClient.AddCommentToPost(*ctx, in)
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
