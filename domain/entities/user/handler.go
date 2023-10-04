package user

import (
	handlerPort "go-port-and-adapter/ports/domain"
	repositoryPort "go-port-and-adapter/ports/repository"
)

type (
	userHandler struct {
		userRepository repositoryPort.IUserRepository
	}
)

func New(userRepository repositoryPort.IUserRepository) handlerPort.IUserHandler {
	return &userHandler{
		userRepository,
	}
}
