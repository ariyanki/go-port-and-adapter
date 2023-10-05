package auth

import (
	"time"

	handlerDto "go-port-and-adapter/ports/domain/dto"
	domainError "go-port-and-adapter/ports/domain/constants/error"
	repoError "go-port-and-adapter/ports/repository/constants/error"
	"go-port-and-adapter/ports/repository/constants/user_status"
	"go-port-and-adapter/helpers/config"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/spf13/viper"
	"golang.org/x/crypto/bcrypt"
)

func (s *authHandler) SignIn(signInRequest handlerDto.SignInRequest) (handlerDto.SignInResponse, error) {
	user, err := s.userRepository.GetUserByUsername(signInRequest.Username)
	if err == repoError.RecordNotFound {
		return handlerDto.SignInResponse{}, domainError.InvalidUserNameOrPassword
	}
	
	if user.Status != user_status.Active {
		return handlerDto.SignInResponse{}, domainError.UserNotActive
	}
	
	if !CheckPasswordHash(user.Password, user.Password) {
		return handlerDto.SignInResponse{}, domainError.InvalidUserNameOrPassword
	}

	// Set custom claims
	claims := &config.JwtCustomClaims{
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

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 8)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}