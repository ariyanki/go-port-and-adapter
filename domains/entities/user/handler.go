package user

import (
	handlerPort "go-port-and-adapter/ports/domain"
	repositoryPort "go-port-and-adapter/ports/repository"
	"go-port-and-adapter/ports/storage"
)

type (
	userHandler struct {
		userRepository repositoryPort.IUserRepository
		storageManager storage.IStorageManager
	}
)

func New(userRepository repositoryPort.IUserRepository, storageManager storage.IStorageManager) handlerPort.IUserHandler {
	return &userHandler{
		userRepository,
		storageManager,
	}
}
