package user

import (
	handlerDto "go-port-and-adapter/ports/domain/dto"
)

func (s *userHandler) GetUserById(ID string) (handlerDto.UserDto, error) {
	getUserById, err := s.userRepository.GetUserByID(ID)
	user := handlerDto.UserDto{
		ID:        getUserById.ID,
		Username:  getUserById.Username,
		Password:  getUserById.Password,
		CreatedAt: getUserById.CreatedAt,
		UpdatedAt: getUserById.CreatedAt,
		DeletedAt: getUserById.DeletedAt,
	}
	return user, err
}