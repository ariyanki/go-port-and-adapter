package auth

import (
	handlerPort "go-port-and-adapter/ports/domain"
)

type AuthEndpoint struct {
	authHandler handlerPort.IAuthHandler
}

func New(authHandler handlerPort.IAuthHandler) *AuthEndpoint {
	return &AuthEndpoint{
		authHandler,
	}
}

