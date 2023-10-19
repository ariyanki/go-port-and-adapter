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

func TestListUser(t *testing.T) {
	repository := repositoryMock.NewUserRepository()
	repository.On("ListUser", mock.Anything)
	repository.On("GetUserById", mock.Anything)
	storage := storageMock.NewStorageManager()

	t.Run("Expect List` User Success", func(t *testing.T) {
		userHandler := userHandler.New(repository, storage)
		var users []handlerDto.UserDto
		err := userHandler.ListUser(handlerDto.ListRequestDto{}, &users)
		assert.Nil(t, err)
	})
}
