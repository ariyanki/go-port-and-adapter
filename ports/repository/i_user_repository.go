package repository

import (
	"go-port-and-adapter/ports/repository/dto"
)

// IUserRepository
type IUserRepository interface {

	// GetUserByUsername
	GetUserByUsername(username string, user *dto.UserDto) error

	// CreateUser
	CreateUser(user dto.CreateUserDto) error

	// GetUserById
	GetUserById(id int, user *dto.UserDto) error

	// UpdateUser
	UpdateUser(user dto.UpdateUserDto) error

	// DeleteUser
	DeleteUser(id int) error

	// ListUser
	ListUser(listRequest dto.ListRequestDto, users *[]dto.UserDto) error
}
