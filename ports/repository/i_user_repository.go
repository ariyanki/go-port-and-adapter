package repository

import (
	"go-port-and-adapter/ports/repository/dto"
)

// Repository is outbound port
type IUserRepository interface {
	//CreateUser insert new data
	CreateUser(user dto.UserDto) error

	//GetUserByID get data by ID
	GetUserByID(ID string) (dto.UserDto, error)

	//GetUserByUsername get data by Username
	GetUserByUsername(username string) (dto.UserDto, error)
}
