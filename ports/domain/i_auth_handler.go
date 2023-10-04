package handler

import (
	"go-port-and-adapter/ports/domain/dto"
)

// Handler is inbound port
type IAuthHandler interface {
	//SignIn signin process
	SignIn(signInRequest dto.SignInRequest) (dto.SignInResponse, error)
}
