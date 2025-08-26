package resource

import (
	"github.com/relaunch-cot/bff-relaunch/config"
	"github.com/relaunch-cot/bff-relaunch/grpc"
	"github.com/relaunch-cot/bff-relaunch/handler"
	"github.com/relaunch-cot/bff-relaunch/server"
	pbUser "github.com/relaunch-cot/lib-relaunch-cot/proto/user"
)

var Servers server.Servers
var Handlers handler.Handlers
var Grpc grpc.Grpc

func Inject() {
	openUserGrpcConnection := openGrpcClientConn(config.USER_MICROSEVICE_CONN, pbUser.NewUserServiceClient)

	Grpc.InjectUser(openUserGrpcConnection)

	Handlers.Inject(&Grpc)
	Servers.Inject(&Handlers)
}
