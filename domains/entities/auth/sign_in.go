package auth

import (
	domainError "go-port-and-adapter/ports/domain/constants/error"
	handlerDto "go-port-and-adapter/ports/domain/dto"
	repoError "go-port-and-adapter/ports/repository/constants/error"
	"go-port-and-adapter/ports/repository/constants/user_status"
	"go-port-and-adapter/systems"
)

func (s *authHandler) SignIn(signInRequest handlerDto.SignInRequest) (handlerDto.SignInResponse, error) {
	user, err := s.userRepository.GetUserByUsername(signInRequest.Username)
	if err == repoError.RecordNotFound {
		return handlerDto.SignInResponse{}, domainError.InvalidUserNameOrPassword
	}

	if user.Status != user_status.Active {
		return handlerDto.SignInResponse{}, domainError.UserNotActive
	}

	if systems.GeneratePassword(signInRequest.Username, signInRequest.Password, user.PasswordSalt) != user.Password {
		return handlerDto.SignInResponse{}, domainError.InvalidUserNameOrPassword
	}
	return handlerDto.SignInResponse{
		ID: user.ID,
	}, err
}
