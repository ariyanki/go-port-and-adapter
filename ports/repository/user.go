package repository

import (
	"go-port-and-adapter/ports/repository/dto"
)

// Repository is outbound port
type UserRepository interface {
	//CreateUser insert new data
	CreateUser(user dto.CreateUserDto) error

	//ReadData get data by ID
	ReadData(ID string) (dto.UserDto, error)
}
