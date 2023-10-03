package user

import (
	domainPort "go-port-and-adapter/ports/domain"
)

type UserController struct {
	userDomain domainPort.UserDomain
}

func New(userDomain domainPort.UserDomain) *UserController {
	return &UserController{
		userDomain,
	}
}

