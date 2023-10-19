package auth

import (
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
	"github.com/spf13/viper"

	"go-port-and-adapter/apps/http/api/v1/middleware"
	handlerDto "go-port-and-adapter/ports/domain/dto"
	"go-port-and-adapter/systems/validator"
)

// @Summary      Sign In
// @Description  Sign In Account
// @Tags         auth
// @Param request body auth.SignInRequest true "SignIn Request Body"
// @Success 200 {object} auth.SignInResponse
// @Router /auth/signin [post]
func (endpoint *AuthEndpoint) SignIn(c echo.Context) error {
	signInRequest := new(SignInRequest)

	if err := c.Bind(signInRequest); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, err)
	}

	if err := validator.GetValidator().Struct(signInRequest); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	signInHandlerRequest := handlerDto.SignInRequest{
		Username: signInRequest.Username,
		Password: signInRequest.Password,
	}
	var signInHandlerResponse handlerDto.SignInResponse
	if err := endpoint.authHandler.SignIn(signInHandlerRequest, &signInHandlerResponse); err != nil {
		return c.JSON(http.StatusUnauthorized, err.Error())
	}

	// Set custom claims
	claims := &middleware.JwtCustomClaims{
		signInHandlerResponse.ID,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 72)),
		},
	}

	// Create token with claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Generate encoded token and send it as response.
	t, err := token.SignedString([]byte(viper.GetString("jwtSign")))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	signInResponse := SignInResponse{
		Token: t,
	}

	return c.JSON(http.StatusOK, signInResponse)
}

type SignInRequest struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type SignInResponse struct {
	Token string `json:"token"`
}
