package repository

import (
	"go-port-and-adapter/ports/repository/dto"
)

// Repository is outbound port (Interface)
type IUserRepository interface {
	//CreateUser insert new data
	CreateUser(user dto.CreateUserDto) error

	//ReadData get data by ID
	ReadData(ID string) (dto.UserDto, error)
}
