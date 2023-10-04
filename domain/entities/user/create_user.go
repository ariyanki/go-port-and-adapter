package user

import (
	repoDto "go-port-and-adapter/ports/repository/dto"
	handlerDto "go-port-and-adapter/ports/domain/dto"
)

func (s *userHandler) CreateUser(user handlerDto.CreateUserDto) error {
	saveUser := repoDto.UserDto{
		Username:  user.Username,
		Password: user.Password,
	}
	return s.userRepository.CreateUser(saveUser)
}