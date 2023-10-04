package auth

import (
	handlerDto "go-port-and-adapter/ports/domain/dto"
)

func (s *authHandler) SignIn(signInRequest handlerDto.SignInRequest) (handlerDto.SignInResponse, error) {
	user, err := s.userRepository.GetUserByUsername(signInRequest.Username)
	response := handlerDto.SignInResponse {
		Username:  user.Username,
		Password:  user.Password,
	}
	return response, err
}