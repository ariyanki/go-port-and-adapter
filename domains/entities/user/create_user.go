package user

import (
	handlerDto "go-port-and-adapter/ports/domain/dto"
	repositoryDto "go-port-and-adapter/ports/repository/dto"
	storageDto "go-port-and-adapter/ports/storage/dto"
	"go-port-and-adapter/systems/automap"
)

func (s *userHandler) CreateUser(request handlerDto.CreateUserDto) error {
	// map request domain to request repository
	var user repositoryDto.CreateUserDto
	if err := automap.AutoMap(request, &user); err != nil {
		return err
	}

	if err := s.storageManager.StoreUserPhoto(storageDto.UserPhotoDto{PhotoFilename: request.PhotoFilename, PhotoFiledata: request.PhotoFiledata}); err != nil {
		return err
	}

	return s.userRepository.CreateUser(user)
}
