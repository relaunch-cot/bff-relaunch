package server

import (
	"github.com/relaunch-cot/bff/handler"
)

type Servers struct {
	User IUser
}

type resource struct {
	handler *handler.Handlers
}

func (c *Servers) Inject(handler *handler.Handlers) {
	c.User = NewUserServer(handler)
}
