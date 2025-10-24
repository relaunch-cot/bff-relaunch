package server

import (
	"github.com/relaunch-cot/bff-relaunch/handler"
)

type Servers struct {
	User         IUser
	Chat         IChat
	Project      IProject
	Notification INotification
}

type resource struct {
	handler *handler.Handlers
}

func (c *Servers) Inject(handler *handler.Handlers) {
	c.User = NewUserServer(handler)
	c.Chat = NewChatServer(handler)
	c.Project = NewProjectServer(handler)
	c.Notification = NewNotificationServer(handler)
}
