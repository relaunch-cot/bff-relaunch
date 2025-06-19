package handler

import handler "bff.com/m/handler/user"

type Handlers struct {
	User IUser
}

func (c *Handlers) Inject(rgrpc *grpc.GRPC) {
	c.User = handler.NewUserHandler(grpc)
}
