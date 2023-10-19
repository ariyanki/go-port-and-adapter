package auth

import (
	domainError "go-port-and-adapter/ports/domain/constants/error"
	handlerDto "go-port-and-adapter/ports/domain/dto"
	"go-port-and-adapter/ports/repository/constants/user_status"
	repositoryDto "go-port-and-adapter/ports/repository/dto"
	"go-port-and-adapter/systems"
)

func (s *authHandler) SignIn(signInRequest handlerDto.SignInRequest, signInResponse *handlerDto.SignInResponse) error {
	var user repositoryDto.UserDto
	if err := s.userRepository.GetUserByUsername(signInRequest.Username, &user); err != nil {
		return err
	}

	if user.Status != user_status.Active {
		return domainError.UserNotActive
	}

	if systems.GeneratePassword(signInRequest.Username, signInRequest.Password, user.PasswordSalt) != user.Password {
		return domainError.InvalidUserNameOrPassword
	}
	signInResponse.ID = user.ID
	return nil
}
