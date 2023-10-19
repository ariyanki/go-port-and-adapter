package user

import (
	handlerDto "go-port-and-adapter/ports/domain/dto"
	repositoryDto "go-port-and-adapter/ports/repository/dto"
	"go-port-and-adapter/systems/automap"
)

func (s *userHandler) GetUserById(id int, response *handlerDto.UserDto) error {
	var user repositoryDto.UserDto
	if err := s.userRepository.GetUserById(id, &user); err != nil {
		return err
	}

	if err := automap.AutoMap(user, &response); err != nil {
		return err
	}

	if err := s.userRepository.GetUserById(user.CreatedBy, &user); err != nil {
		return err
	}
	response.CreatedByFullname = user.Fullname

	if err := s.userRepository.GetUserById(user.UpdatedBy, &user); err != nil {
		return err
	}
	response.UpdatedByFullname = user.Fullname
	return nil
}
