package middleware

import (
	"net/http"

	"github.com/golang-jwt/jwt/v5"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/spf13/viper"
)

// JwtCustomClaims JwtCustomClaims
type JwtCustomClaims struct {
	ID int `json:"id"`
	jwt.RegisteredClaims
}

// JwtConfig JwtConfig
var JwtConfig echojwt.Config

func init() {
	JwtConfig = echojwt.Config{
		NewClaimsFunc: func(c echo.Context) jwt.Claims {
			return new(JwtCustomClaims)
		},
		SigningKey: []byte(viper.GetString("jwtSign")),
		ErrorHandler: func(c echo.Context, err error) error {
			if err != nil {
				return c.JSON(http.StatusUnauthorized, "Unauthorized: "+err.Error())
			}
			return nil
		},
	}
}
