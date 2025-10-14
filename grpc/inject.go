package grpc

import (
	"github.com/relaunch-cot/bff-relaunch/grpc/chat"
	"github.com/relaunch-cot/bff-relaunch/grpc/user"
	pbChat "github.com/relaunch-cot/lib-relaunch-cot/proto/chat"
	pbUser "github.com/relaunch-cot/lib-relaunch-cot/proto/user"
)

type Grpc struct {
	UserGRPC user.IUserGRPC
	ChatGRPC chat.IChatGRPC
}

func (g *Grpc) Inject(
	grpcUser pbUser.UserServiceClient,
	grpcChat pbChat.ChatServiceClient,
) {
	g.UserGRPC = user.NewUserGrpcClient(grpcUser)
	g.ChatGRPC = chat.NewChatGrpcClient(grpcChat)
}
