package auth

import (
	"net/http"

	"github.com/labstack/echo/v4"
	
	handlerDto "go-port-and-adapter/ports/domain/dto"
)

// @Summary      Sign In
// @Description  Sign In Account
// @Tags         auth
// @Param request body auth.SignInRequest true "SignIn Request Body"
// @Success 200 {object} auth.SignInResponse
// @Router /auth/signin [post]
func (endpoint *AuthEndpoint) SignIn(c echo.Context) error {
	signInRequest := new(SignInRequest)

	signInHandlerRequest := handlerDto.SignInRequest{
		Username:  signInRequest.Username,
		Password: signInRequest.Password,
	}

	signInHandlerResponse, err := endpoint.authHandler.SignIn(signInHandlerRequest)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, err.Error())
	}

	signInResponse := SignInResponse{
		Token:  signInHandlerResponse.Token,
	}

	return c.JSON(http.StatusOK, signInResponse)
}

type SignInRequest struct {
	Username  string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type SignInResponse struct {
	Token  string `json:"token"`
}