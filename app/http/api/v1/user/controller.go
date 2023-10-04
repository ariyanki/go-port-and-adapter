package user

import (
	domainPort "go-port-and-adapter/ports/domain"
)

type UserController struct {
	userDomain domainPort.IUserDomain
}

func New(userDomain domainPort.IUserDomain) *UserController {
	return &UserController{
		userDomain,
	}
}

