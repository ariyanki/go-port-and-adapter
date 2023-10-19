package user_test

import (
	"testing"

	userHandler "go-port-and-adapter/domains/entities/user"
	repositoryMock "go-port-and-adapter/tests/unit_tests/mocks/adapters/repository/mysql"
	storageMock "go-port-and-adapter/tests/unit_tests/mocks/adapters/storage"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestDeleteUser(t *testing.T) {
	repository := repositoryMock.NewUserRepository()
	repository.On("DeleteUser", mock.Anything)
	storage := storageMock.NewStorageManager()

	t.Run("Expect Delete User Success", func(t *testing.T) {
		userHandler := userHandler.New(repository, storage)
		err := userHandler.DeleteUser(1)
		assert.Nil(t, err)
	})

	t.Run("Expect Delete User not Found", func(t *testing.T) {
		userHandler := userHandler.New(repository, storage)
		err := userHandler.DeleteUser(123)
		assert.NotNil(t, err)
	})
}
