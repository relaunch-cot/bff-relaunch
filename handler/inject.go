package handler

import "github.com/relaunch-cot/bff-relaunch/grpc"

type Handlers struct {
	User         IUser
	Chat         IChat
	Project      IProject
	Notification INotification
	Post         IPost
}

func (c *Handlers) Inject(grpc *grpc.Grpc) {
	c.User = NewUserHandler(grpc)
	c.Chat = NewChatHandler(grpc)
	c.Project = NewProjectHandler(grpc)
	c.Notification = NewNotificationHandler(grpc)
	c.Post = NewPostHandler(grpc)
}
