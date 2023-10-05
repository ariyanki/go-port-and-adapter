package auth

import (
	"time"
	"crypto/sha256"
	"encoding/hex"

	handlerDto "go-port-and-adapter/ports/domain/dto"
	domainError "go-port-and-adapter/ports/domain/constants/error"
	repoError "go-port-and-adapter/ports/repository/constants/error"
	"go-port-and-adapter/ports/repository/constants/user_status"
	"go-port-and-adapter/helpers/config"

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

	if GeneratePassword(signInRequest.Username, signInRequest.Password, user.PasswordSalt)!=user.Password {
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

func GeneratePassword(username, password, salt string) string {
	input := username + password + salt

    // Create a new SHA-256 hash object
    hasher := sha256.New()

    // Write the string to the hash object
    _, err := hasher.Write([]byte(input))
    if err != nil {
        return ""
    }

    // Get the hash sum as a byte slice
    hashBytes := hasher.Sum(nil)

    // Convert the hash bytes to a hexadecimal string
    hashHex := hex.EncodeToString(hashBytes)
	return hashHex
}