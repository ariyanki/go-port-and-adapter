package domain

import (
	"go-port-and-adapter/ports/domain/dto"
)

// Domain is inbound port
type IUserDomain interface {
	//CreateUser insert new data
	CreateUser(user dto.CreateUserDto) error

	//ReadData get data by ID
	ReadData(ID string) (dto.UserDto, error)
}
