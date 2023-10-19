package user

import (
	handlerPort "go-port-and-adapter/ports/domain"
)

type UserEndpoint struct {
	userHandler handlerPort.IUserHandler
}

func New(userHandler handlerPort.IUserHandler) *UserEndpoint {
	return &UserEndpoint{
		userHandler,
	}
}
