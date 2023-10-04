package handler

import (
	"go-port-and-adapter/ports/domain/dto"
)

// Handler is inbound port
type IUserHandler interface {
	//CreateUser insert new data
	CreateUser(user dto.CreateUserDto) error

	//ReadData get data by ID
	ReadData(ID string) (dto.UserDto, error)
}
