package auth

import (
	handlerPort "go-port-and-adapter/ports/domain"
	repositoryPort "go-port-and-adapter/ports/repository"
)

type (
	authHandler struct {
		userRepository repositoryPort.IUserRepository
	}
)

func New(userRepository repositoryPort.IUserRepository) handlerPort.IAuthHandler {
	return &authHandler{
		userRepository,
	}
}
