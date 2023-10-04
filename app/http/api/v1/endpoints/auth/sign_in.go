package auth

import (
	"net/http"

	"github.com/labstack/echo/v4"
	
	handlerDto "go-port-and-adapter/ports/domain/dto"
)

func (endpoint *AuthEndpoint) SignIn(c echo.Context) error {
	apiRequest := new(SignInRequest)

	signInRequest := handlerDto.SignInRequest{
		Username:  apiRequest.Username,
		Password: apiRequest.Password,
	}

	signInResponse, err := endpoint.authHandler.SignIn(signInRequest)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, err)
	}

	apiResponse := SignInResponse{
		Username:  signInResponse.Username,
		Password: signInResponse.Password,
	}

	return c.JSON(http.StatusOK, apiResponse)
}

type SignInRequest struct {
	Username  string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type SignInResponse struct {
	Username  string `json:"username"`
	Password string `json:"password"`
}
