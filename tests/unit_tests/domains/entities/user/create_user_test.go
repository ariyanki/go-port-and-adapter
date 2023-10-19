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

func TestCreateUser(t *testing.T) {
	repository := repositoryMock.NewUserRepository()
	repository.On("CreateUser", mock.Anything)
	storage := storageMock.NewStorageManager()

	t.Run("Expect Create User Success", func(t *testing.T) {
		userHandler := userHandler.New(repository, storage)
		err := userHandler.CreateUser(handlerDto.CreateUserDto{})
		assert.Nil(t, err)
	})

	t.Run("Expect Create User Failed", func(t *testing.T) {
		userHandler := userHandler.New(repository, storage)
		err := userHandler.CreateUser(handlerDto.CreateUserDto{Username: "failed"})
		assert.NotNil(t, err)
	})
}
