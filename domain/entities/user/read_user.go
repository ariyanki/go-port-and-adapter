package user

import (
	handlerDto "go-port-and-adapter/ports/domain/dto"
)

func (s *userHandler) ReadData(ID string) (handlerDto.UserDto, error) {
	readData, err := s.userRepository.GetUserByID(ID)
	user := handlerDto.UserDto{
		ID:        readData.ID,
		Username:  readData.Username,
		Password:  readData.Password,
		CreatedAt: readData.CreatedAt,
		UpdatedAt: readData.CreatedAt,
		DeletedAt: readData.DeletedAt,
	}
	return user, err
}