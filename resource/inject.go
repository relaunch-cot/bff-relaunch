package resource

import (
	"github.com/relaunch-cot/bff-relaunch/config"
	"github.com/relaunch-cot/bff-relaunch/grpc"
	"github.com/relaunch-cot/bff-relaunch/handler"
	"github.com/relaunch-cot/bff-relaunch/server"
	pbChat "github.com/relaunch-cot/lib-relaunch-cot/proto/chat"
	pbNotification "github.com/relaunch-cot/lib-relaunch-cot/proto/notification"
	pbPost "github.com/relaunch-cot/lib-relaunch-cot/proto/post"
	pbProject "github.com/relaunch-cot/lib-relaunch-cot/proto/project"
	pbUser "github.com/relaunch-cot/lib-relaunch-cot/proto/user"
)

var Servers server.Servers
var Handlers handler.Handlers
var Grpc grpc.Grpc

func Inject() {
	openUserGrpcConnection := openGrpcClientConn(config.USER_MICROSERVICE_CONN, pbUser.NewUserServiceClient)
	openChatGrpcConnection := openGrpcClientConn(config.CHAT_MICROSSEVICE_CONN, pbChat.NewChatServiceClient)
	openProjectGrpcConnection := openGrpcClientConn(config.PROJECT_MICROSSERVICE_CONN, pbProject.NewProjectServiceClient)
	openNotificationGrpcConnection := openGrpcClientConn(config.NOTIFICATION_MICROSERVICE_CONN, pbNotification.NewNotificationServiceClient)
	openPostGrpcConnection := openGrpcClientConn(config.POST_MICROSERVICE_CONN, pbPost.NewPostServiceClient)

	Grpc.Inject(openUserGrpcConnection, openChatGrpcConnection, openProjectGrpcConnection, openNotificationGrpcConnection, openPostGrpcConnection)

	Handlers.Inject(&Grpc)
	Servers.Inject(&Handlers)
}
