package handler

import "github.com/relaunch-cot/bff/grpc"

type Handlers struct {
	User IUser
}

func (c *Handlers) Inject(grpc *grpc.Grpc) {
	c.User = NewUserHandler(grpc)
}
