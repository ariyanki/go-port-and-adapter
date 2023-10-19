package user_test

import (
	"testing"

	userHandler "go-port-and-adapter/domains/entities/user"
	handlerDto "go-port-and-adapter/ports/domain/dto"
	repositoryMock "go-port-and-adapter/tests/unit_tests/mocks/adapters/repository/mysql"
	storageMock "go-port-and-adapter/tests/unit_tests/mocks/adapters/storage"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestGetUserById(t *testing.T) {
	repository := repositoryMock.NewUserRepository()
	repository.On("GetUserById", mock.Anything)
	storage := storageMock.NewStorageManager()

	t.Run("Expect Get User Success", func(t *testing.T) {
		userHandler := userHandler.New(repository, storage)
		err := userHandler.GetUserById(1, &handlerDto.UserDto{})
		assert.Nil(t, err)
	})

	t.Run("Expect Get User not Found", func(t *testing.T) {
		userHandler := userHandler.New(repository, storage)
		err := userHandler.GetUserById(123, &handlerDto.UserDto{})
		assert.NotNil(t, err)
	})
}
