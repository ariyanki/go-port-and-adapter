package auth

import (
	"time"

	handlerDto "go-port-and-adapter/ports/domain/dto"
	systemDto "go-port-and-adapter/ports/system/dto"
	domainError "go-port-and-adapter/ports/domain/constants/error"
	repoError "go-port-and-adapter/ports/repository/constants/error"
	"go-port-and-adapter/ports/repository/constants/user_status"
	"go-port-and-adapter/systems"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/spf13/viper"
)

func (s *authHandler) SignIn(signInRequest handlerDto.SignInRequest) (handlerDto.SignInResponse, error) {
	user, err := s.userRepository.GetUserByUsername(signInRequest.Username)
	if err == repoError.RecordNotFound {
		return handlerDto.SignInResponse{}, domainError.InvalidUserNameOrPassword
	}
	
	if user.Status != user_status.Active {
		return handlerDto.SignInResponse{}, domainError.UserNotActive
	}

	if systems.GeneratePassword(signInRequest.Username, signInRequest.Password, user.PasswordSalt)!=user.Password {
		return handlerDto.SignInResponse{}, domainError.InvalidUserNameOrPassword
	}

	// Set custom claims
	claims := &systemDto.JwtCustomClaims{
		user.ID,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 72).Unix(),
		},
	}

	// Create token with claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Generate encoded token and send it as response.
	t, err := token.SignedString([]byte(viper.GetString("jwtSign")))
	if err != nil {
		return handlerDto.SignInResponse{}, err
	}

	response := handlerDto.SignInResponse {
		Token:  t,
	}
	return response, err
}