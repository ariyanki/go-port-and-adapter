package handler

import (
	"go-port-and-adapter/ports/domain/dto"
)

// IAuthHandler
type IAuthHandler interface {
	// SignIn
	SignIn(signInRequest dto.SignInRequest, signInResponse *dto.SignInResponse) error
}
