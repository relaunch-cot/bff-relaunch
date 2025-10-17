package grpc

import (
	"github.com/relaunch-cot/bff-relaunch/grpc/chat"
	"github.com/relaunch-cot/bff-relaunch/grpc/project"
	"github.com/relaunch-cot/bff-relaunch/grpc/user"
	pbChat "github.com/relaunch-cot/lib-relaunch-cot/proto/chat"
	pbProject "github.com/relaunch-cot/lib-relaunch-cot/proto/project"
	pbUser "github.com/relaunch-cot/lib-relaunch-cot/proto/user"
)

type Grpc struct {
	UserGRPC    user.IUserGRPC
	ChatGRPC    chat.IChatGRPC
	ProjectGRPC project.IProjectGRPC
}

func (g *Grpc) Inject(
	grpcUser pbUser.UserServiceClient,
	grpcChat pbChat.ChatServiceClient,
	grpcProject pbProject.ProjectServiceClient,
) {
	g.UserGRPC = user.NewUserGrpcClient(grpcUser)
	g.ChatGRPC = chat.NewChatGrpcClient(grpcChat)
	g.ProjectGRPC = project.NewProjectGrpcClient(grpcProject)
}
