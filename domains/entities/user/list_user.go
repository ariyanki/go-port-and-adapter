package user

import (
	handlerDto "go-port-and-adapter/ports/domain/dto"
	repositoryDto "go-port-and-adapter/ports/repository/dto"
	"go-port-and-adapter/systems/automap"
)

func (s *userHandler) ListUser(request handlerDto.ListRequestDto, response *[]handlerDto.UserDto) error {
	// map request domain to request repository
	var listRequestDto repositoryDto.ListRequestDto
	if err := automap.AutoMap(request, &listRequestDto); err != nil {
		return err
	}

	// Get List Data
	var usersDto []repositoryDto.UserDto
	if err := s.userRepository.ListUser(listRequestDto, &usersDto); err != nil {
		return err
	}

	// map response repository to response domain
	if err := automap.AutoMap(usersDto, &response); err != nil {
		return err
	}
	var user repositoryDto.UserDto
	for _, w := range *response {
		s.userRepository.GetUserById(w.CreatedBy, &user)
		w.CreatedByFullname = user.Fullname
		s.userRepository.GetUserById(w.UpdatedBy, &user)
		w.CreatedByFullname = user.Fullname
	}
	return nil
}
