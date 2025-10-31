package handler

import (
	"context"

	"github.com/relaunch-cot/bff-relaunch/grpc"
	pb "github.com/relaunch-cot/lib-relaunch-cot/proto/chat"
)

type IChat interface {
	CreateNewChat(ctx *context.Context, in *pb.CreateNewChatRequest) error
	SendMessage(ctx *context.Context, in *pb.SendMessageRequest) error
	GetAllMessagesFromChat(ctx *context.Context, in *pb.GetAllMessagesFromChatRequest) (*pb.GetAllMessagesFromChatResponse, error)
	GetAllChatsFromUser(ctx *context.Context, in *pb.GetAllChatsFromUserRequest) (*pb.GetAllChatsFromUserResponse, error)
	GetChatFromUsers(ctx *context.Context, in *pb.GetChatFromUsersRequest) (*pb.GetChatFromUsersResponse, error)
	GetChatById(ctx *context.Context, in *pb.GetChatByIdRequest) (*pb.GetChatByIdResponse, error)
}

type chatResource struct {
	grpc *grpc.Grpc
}

func (r *chatResource) CreateNewChat(ctx *context.Context, in *pb.CreateNewChatRequest) error {
	err := r.grpc.ChatGRPC.CreateNewChat(ctx, in)
	if err != nil {
		return err
	}

	return nil
}

func (r *chatResource) SendMessage(ctx *context.Context, in *pb.SendMessageRequest) error {
	err := r.grpc.ChatGRPC.SendMessage(ctx, in)
	if err != nil {
		return err
	}

	return nil
}

func (r *chatResource) GetAllMessagesFromChat(ctx *context.Context, in *pb.GetAllMessagesFromChatRequest) (*pb.GetAllMessagesFromChatResponse, error) {
	response, err := r.grpc.ChatGRPC.GetAllMessagesFromChat(ctx, in)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (r *chatResource) GetAllChatsFromUser(ctx *context.Context, in *pb.GetAllChatsFromUserRequest) (*pb.GetAllChatsFromUserResponse, error) {
	response, err := r.grpc.ChatGRPC.GetAllChatsFromUser(ctx, in)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (r *chatResource) GetChatFromUsers(ctx *context.Context, in *pb.GetChatFromUsersRequest) (*pb.GetChatFromUsersResponse, error) {
	response, err := r.grpc.ChatGRPC.GetChatFromUsers(ctx, in)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (r *chatResource) GetChatById(ctx *context.Context, in *pb.GetChatByIdRequest) (*pb.GetChatByIdResponse, error) {
	response, err := r.grpc.ChatGRPC.GetChatById(ctx, in)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func NewChatHandler(grpc *grpc.Grpc) IChat {
	return &chatResource{
		grpc: grpc,
	}
}
