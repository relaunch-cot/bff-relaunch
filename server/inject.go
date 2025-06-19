package server

type Servers struct {
	User IUser
}

type resource struct {
	handler *handler.Handlers
}
