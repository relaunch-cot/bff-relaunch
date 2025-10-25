package chat

import (
	"context"

	pb "github.com/relaunch-cot/lib-relaunch-cot/proto/chat"
)

type IChatGRPC interface {
	CreateNewChat(ctx *context.Context, in *pb.CreateNewChatRequest) error
	SendMessage(ctx *context.Context, in *pb.SendMessageRequest) error
	GetAllMessagesFromChat(ctx *context.Context, in *pb.GetAllMessagesFromChatRequest) (*pb.GetAllMessagesFromChatResponse, error)
	GetAllChatsFromUser(ctx *context.Context, in *pb.GetAllChatsFromUserRequest) (*pb.GetAllChatsFromUserResponse, error)
	GetChatFromUsers(ctx *context.Context, in *pb.GetChatFromUsersRequest) (*pb.GetChatFromUsersResponse, error)
}

type resource struct {
	grpcClient pb.ChatServiceClient
}

func (r *resource) CreateNewChat(ctx *context.Context, in *pb.CreateNewChatRequest) error {
	_, err := r.grpcClient.CreateNewChat(*ctx, in)
	if err != nil {
		return err
	}

	return nil
}

func (r *resource) SendMessage(ctx *context.Context, in *pb.SendMessageRequest) error {
	_, err := r.grpcClient.SendMessage(*ctx, in)
	if err != nil {
		return err
	}

	return nil
}

func (r *resource) GetAllMessagesFromChat(ctx *context.Context, in *pb.GetAllMessagesFromChatRequest) (*pb.GetAllMessagesFromChatResponse, error) {
	response, err := r.grpcClient.GetAllMessagesFromChat(*ctx, in)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (r *resource) GetAllChatsFromUser(ctx *context.Context, in *pb.GetAllChatsFromUserRequest) (*pb.GetAllChatsFromUserResponse, error) {
	response, err := r.grpcClient.GetAllChatsFromUser(*ctx, in)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (r *resource) GetChatFromUsers(ctx *context.Context, in *pb.GetChatFromUsersRequest) (*pb.GetChatFromUsersResponse, error) {
	response, err := r.grpcClient.GetChatFromUsers(*ctx, in)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func NewChatGrpcClient(grpcClient pb.ChatServiceClient) IChatGRPC {
	return &resource{
		grpcClient: grpcClient,
	}
}
