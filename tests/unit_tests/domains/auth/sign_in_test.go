package auth_test

import (
	"testing"

	handlerDto "go-port-and-adapter/ports/domain/dto"
	authHandler "go-port-and-adapter/domains/entities/auth"
	repositoryMock "go-port-and-adapter/tests/unit_tests/mocks/adapters/repository/mysql"
	domainError "go-port-and-adapter/ports/domain/constants/error"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestSignIn(t *testing.T) {
	t.Run("Expect sign in success", func(t *testing.T) {
		repository := repositoryMock.NewUserRepository()
		repository.On("GetUserByUsername", mock.Anything)

		authHandler := authHandler.New(repository)
		_, err := authHandler.SignIn(handlerDto.SignInRequest{
			Username: "admin@admin.com",
			Password: "123456",
		})
		assert.Nil(t, err)
	})

	t.Run("Expect sign in user not active", func(t *testing.T) {
		repository := repositoryMock.NewUserRepository()
		repository.On("GetUserByUsername", mock.Anything)

		authHandler := authHandler.New(repository)
		_, err := authHandler.SignIn(handlerDto.SignInRequest{
			Username: "admin",
			Password: "123456",
		})
		assert.NotNil(t, err)
		assert.Equal(t, err, domainError.UserNotActive)
	})

	t.Run("Expect sign in invalid username or password", func(t *testing.T) {
		repository := repositoryMock.NewUserRepository()
		repository.On("GetUserByUsername", mock.Anything)

		authHandler := authHandler.New(repository)
		_, err := authHandler.SignIn(handlerDto.SignInRequest{
			Username: "admin@admin.com",
			Password: "1234",
		})
		assert.NotNil(t, err)
		assert.Equal(t, err, domainError.InvalidUserNameOrPassword)
	})
}