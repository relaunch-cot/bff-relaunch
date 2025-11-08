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
	GetAllLikesFromPost(ctx *context.Context, in *pb.GetAllLikesFromPostRequest) (*pb.GetAllLikesFromPostResponse, error)
	UpdateLikesFromPostOrComment(ctx *context.Context, in *pb.UpdateLikesFromPostOrCommentRequest) (*pb.UpdateLikesFromPostOrCommentResponse, error)
	CreateCommentOrReply(ctx *context.Context, in *pb.CreateCommentOrReplyRequest) (*pb.CreateCommentOrReplyResponse, error)
	DeleteCommentOrReply(ctx *context.Context, in *pb.DeleteCommentOrReplyRequest) (*pb.DeleteCommentOrReplyResponse, error)
	GetAllCommentsFromPost(ctx *context.Context, in *pb.GetAllCommentsFromPostRequest) (*pb.GetAllCommentsFromPostResponse, error)
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

func (r *postResource) GetAllLikesFromPost(ctx *context.Context, in *pb.GetAllLikesFromPostRequest) (*pb.GetAllLikesFromPostResponse, error) {
	getLikesAllFromPostResponse, err := r.grpcClient.PostGRPC.GetAllLikesFromPost(ctx, in)
	if err != nil {
		return nil, err
	}

	return getLikesAllFromPostResponse, nil
}

func (r *postResource) UpdateLikesFromPostOrComment(ctx *context.Context, in *pb.UpdateLikesFromPostOrCommentRequest) (*pb.UpdateLikesFromPostOrCommentResponse, error) {
	updateLikesFromPostOrCommentResponse, err := r.grpcClient.PostGRPC.UpdateLikesFromPostOrComment(ctx, in)
	if err != nil {
		return nil, err
	}

	return updateLikesFromPostOrCommentResponse, nil
}

func (r *postResource) CreateCommentOrReply(ctx *context.Context, in *pb.CreateCommentOrReplyRequest) (*pb.CreateCommentOrReplyResponse, error) {
	createCommentOrReplyResponse, err := r.grpcClient.PostGRPC.CreateCommentOrReply(ctx, in)
	if err != nil {
		return nil, err
	}

	return createCommentOrReplyResponse, nil
}

func (r *postResource) DeleteCommentOrReply(ctx *context.Context, in *pb.DeleteCommentOrReplyRequest) (*pb.DeleteCommentOrReplyResponse, error) {
	deleteCommentOrReplyResponse, err := r.grpcClient.PostGRPC.DeleteCommentOrReply(ctx, in)
	if err != nil {
		return nil, err
	}

	return deleteCommentOrReplyResponse, nil
}

func (r *postResource) GetAllCommentsFromPost(ctx *context.Context, in *pb.GetAllCommentsFromPostRequest) (*pb.GetAllCommentsFromPostResponse, error) {
	getAllCommentsFromPostResponse, err := r.grpcClient.PostGRPC.GetAllCommentsFromPost(ctx, in)
	if err != nil {
		return nil, err
	}

	return getAllCommentsFromPostResponse, nil
}

func NewPostHandler(grpcClient *grpc.Grpc) IPost {
	return &postResource{
		grpcClient: grpcClient,
	}
}
