package handler

import (
	"go-port-and-adapter/ports/domain/dto"
)

// IUserHandler
type IUserHandler interface {
	// CreateUser
	CreateUser(request dto.CreateUserDto) error

	// GetUserById
	GetUserById(id int, response *dto.UserDto) error

	// UpdateUser
	UpdateUser(request dto.UpdateUserDto) error

	// DeleteUser
	DeleteUser(id int) error

	// ListUser
	ListUser(request dto.ListRequestDto, response *[]dto.UserDto) error
}
