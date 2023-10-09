package middleware

import (
	"go-port-and-adapter/systems/config"
	"net/http"

	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/spf13/viper"
)

// JwtConfig JwtConfig
var JwtConfig echojwt.Config

func init() {
	config.Load()
	JwtConfig = echojwt.Config{
		SigningKey: []byte(viper.GetString("jwtSign")),
		ErrorHandler: func(c echo.Context, err error) error {
			if err != nil {
				return c.JSON(http.StatusUnauthorized, "Unauthorized: "+err.Error())
			}
			return nil
		},
	}
}
