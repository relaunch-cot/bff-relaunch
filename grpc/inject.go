package grpc

import (
	"github.com/relaunch-cot/bff-relaunch/grpc/user"
	pb "github.com/relaunch-cot/lib-relaunch-cot/proto/user"
)

type Grpc struct {
	UserGRPC user.IUserGRPC
}

func (g *Grpc) InjectUser(grpcUser pb.UserServiceClient) {
	g.UserGRPC = user.NewUserGrpcClient(grpcUser)
}
