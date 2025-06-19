package resource

import (
	"github.com/relaunch-cot/bff/config"
	"github.com/relaunch-cot/bff/grpc"
	"github.com/relaunch-cot/bff/handler"
	"github.com/relaunch-cot/bff/server"
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
