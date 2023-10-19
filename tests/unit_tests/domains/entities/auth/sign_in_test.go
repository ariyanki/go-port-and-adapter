package auth_test

import (
	"testing"

	authHandler "go-port-and-adapter/domains/entities/auth"
	domainError "go-port-and-adapter/ports/domain/constants/error"
	handlerDto "go-port-and-adapter/ports/domain/dto"
	repositoryMock "go-port-and-adapter/tests/unit_tests/mocks/adapters/repository/mysql"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestSignIn(t *testing.T) {
	t.Run("Expect sign in success", func(t *testing.T) {
		repository := repositoryMock.NewUserRepository()
		repository.On("GetUserByUsername", mock.Anything)

		authHandler := authHandler.New(repository)
		err := authHandler.SignIn(handlerDto.SignInRequest{
			Username: "admin@admin.com",
			Password: "123456",
		}, &handlerDto.SignInResponse{})
		assert.Nil(t, err)
	})

	t.Run("Expect sign in user not active", func(t *testing.T) {
		repository := repositoryMock.NewUserRepository()
		repository.On("GetUserByUsername", mock.Anything)

		authHandler := authHandler.New(repository)
		err := authHandler.SignIn(handlerDto.SignInRequest{
			Username: "admin",
			Password: "123456",
		}, &handlerDto.SignInResponse{})
		assert.NotNil(t, err)
		assert.Equal(t, err, domainError.UserNotActive)
	})

	t.Run("Expect sign in invalid username or password", func(t *testing.T) {
		repository := repositoryMock.NewUserRepository()
		repository.On("GetUserByUsername", mock.Anything)

		authHandler := authHandler.New(repository)
		err := authHandler.SignIn(handlerDto.SignInRequest{
			Username: "admin@admin.com",
			Password: "1234",
		}, &handlerDto.SignInResponse{})
		assert.NotNil(t, err)
		assert.Equal(t, err, domainError.InvalidUserNameOrPassword)
	})
}
