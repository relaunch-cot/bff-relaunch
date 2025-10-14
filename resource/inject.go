package resource

import (
	"github.com/relaunch-cot/bff-relaunch/config"
	"github.com/relaunch-cot/bff-relaunch/grpc"
	"github.com/relaunch-cot/bff-relaunch/handler"
	"github.com/relaunch-cot/bff-relaunch/server"
	pbChat "github.com/relaunch-cot/lib-relaunch-cot/proto/chat"
	pbUser "github.com/relaunch-cot/lib-relaunch-cot/proto/user"
)

var Servers server.Servers
var Handlers handler.Handlers
var Grpc grpc.Grpc

func Inject() {
	openUserGrpcConnection := openGrpcClientConn(config.USER_MICROSERVICE_CONN, pbUser.NewUserServiceClient)
	openChatGrpcConnection := openGrpcClientConn(config.CHAT_MICROSSEVICE_CONN, pbChat.NewChatServiceClient)

	Grpc.Inject(openUserGrpcConnection, openChatGrpcConnection)

	Handlers.Inject(&Grpc)
	Servers.Inject(&Handlers)
}
