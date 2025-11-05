package handler

import (
	"context"

	"github.com/relaunch-cot/bff-relaunch/grpc"
	pb "github.com/relaunch-cot/lib-relaunch-cot/proto/post"
)

type IPost interface {
	CreatePost(ctx *context.Context, in *pb.CreatePostRequest) error
	GetPost(ctx *context.Context, in *pb.GetPostRequest) (*pb.GetPostResponse, error)
	GetAllPosts(ctx *context.Context) (*pb.GetAllPostsResponse, error)
	GetAllPostsFromUser(ctx *context.Context, in *pb.GetAllPostsFromUserRequest) (*pb.GetAllPostsFromUserResponse, error)
	UpdatePost(ctx *context.Context, in *pb.UpdatePostRequest) (*pb.UpdatePostResponse, error)
	DeletePost(ctx *context.Context, in *pb.DeletePostRequest) error
	UpdateLikesFromPost(ctx *context.Context, in *pb.UpdateLikesFromPostRequest) (*pb.UpdateLikesFromPostResponse, error)
	AddCommentToPost(ctx *context.Context, in *pb.AddCommentToPostRequest) (*pb.AddCommentToPostResponse, error)
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

func (r *postResource) GetPost(ctx *context.Context, in *pb.GetPostRequest) (*pb.GetPostResponse, error) {
	getPostResponse, err := r.grpcClient.PostGRPC.GetPost(ctx, in)
	if err != nil {
		return nil, err
	}

	return getPostResponse, nil
}

func (r *postResource) GetAllPosts(ctx *context.Context) (*pb.GetAllPostsResponse, error) {
	getAllPostsResponse, err := r.grpcClient.PostGRPC.GetAllPosts(ctx)
	if err != nil {
		return nil, err
	}

	return getAllPostsResponse, nil
}

func (r *postResource) GetAllPostsFromUser(ctx *context.Context, in *pb.GetAllPostsFromUserRequest) (*pb.GetAllPostsFromUserResponse, error) {
	getAllPostsFromUserResponse, err := r.grpcClient.PostGRPC.GetAllPostsFromUser(ctx, in)
	if err != nil {
		return nil, err
	}

	return getAllPostsFromUserResponse, nil
}

func (r *postResource) UpdatePost(ctx *context.Context, in *pb.UpdatePostRequest) (*pb.UpdatePostResponse, error) {
	updatePostResponse, err := r.grpcClient.PostGRPC.UpdatePost(ctx, in)
	if err != nil {
		return nil, err
	}

	return updatePostResponse, nil
}

func (r *postResource) DeletePost(ctx *context.Context, in *pb.DeletePostRequest) error {
	err := r.grpcClient.PostGRPC.DeletePost(ctx, in)
	if err != nil {
		return err
	}

	return nil
}

func (r *postResource) UpdateLikesFromPost(ctx *context.Context, in *pb.UpdateLikesFromPostRequest) (*pb.UpdateLikesFromPostResponse, error) {
	updateLikesFromPostResponse, err := r.grpcClient.PostGRPC.UpdateLikesFromPost(ctx, in)
	if err != nil {
		return nil, err
	}

	return updateLikesFromPostResponse, nil
}

func (r *postResource) AddCommentToPost(ctx *context.Context, in *pb.AddCommentToPostRequest) (*pb.AddCommentToPostResponse, error) {
	addCommentToPostResponse, err := r.grpcClient.PostGRPC.AddCommentToPost(ctx, in)
	if err != nil {
		return nil, err
	}

	return addCommentToPostResponse, nil
}

func NewPostHandler(grpcClient *grpc.Grpc) IPost {
	return &postResource{
		grpcClient: grpcClient,
	}
}
