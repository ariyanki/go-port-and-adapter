package handler

import (
	"go-port-and-adapter/ports/domain/dto"
)

// Handler is inbound port
type IUserHandler interface {
	//CreateUser insert new data
	CreateUser(user dto.CreateUserDto) error

	//GetUserById get data by ID
	GetUserById(ID string) (dto.UserDto, error)
}
