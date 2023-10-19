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

func TestUpdateUser(t *testing.T) {
	repository := repositoryMock.NewUserRepository()
	repository.On("UpdateUser", mock.Anything)
	storage := storageMock.NewStorageManager()

	t.Run("Expect Update User Success", func(t *testing.T) {
		userHandler := userHandler.New(repository, storage)
		err := userHandler.UpdateUser(handlerDto.UpdateUserDto{})
		assert.Nil(t, err)
	})

	t.Run("Expect Update User Failed", func(t *testing.T) {
		userHandler := userHandler.New(repository, storage)
		err := userHandler.UpdateUser(handlerDto.UpdateUserDto{Username: "failed"})
		assert.NotNil(t, err)
	})
}
