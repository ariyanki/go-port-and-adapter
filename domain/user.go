package domain

import (
	domainDto "go-port-and-adapter/ports/domain/dto"
	domainPort "go-port-and-adapter/ports/domain"
	repoDto "go-port-and-adapter/ports/repository/dto"
	repositoryPort "go-port-and-adapter/ports/repository"
	
)

type (
	userDomain struct {
		userRepository repositoryPort.UserRepository
	}
)

func NewUserDomain(userRepository repositoryPort.UserRepository) domainPort.UserDomain {
	return &userDomain{
		userRepository,
	}
}

func (s *userDomain) CreateUser(user domainDto.CreateUserDto) error {
	saveUser := repoDto.CreateUserDto{
		Username:  user.Username,
		Password: user.Password,
	}
	return s.userRepository.CreateUser(saveUser)
}

func (s *userDomain) ReadData(ID string) (domainDto.UserDto, error) {
	readData, err := s.userRepository.ReadData(ID)
	user := domainDto.UserDto{
		ID:        readData.ID,
		Username:  readData.Username,
		Password:  readData.Password,
		CreatedAt: readData.CreatedAt,
		UpdatedAt: readData.CreatedAt,
		DeletedAt: readData.DeletedAt,
	}
	return user, err
}
