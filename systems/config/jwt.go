package config

import (
	"net/http"

	"github.com/labstack/echo/v4/middleware"
	"github.com/spf13/viper"
	"go-port-and-adapter/ports/system/dto"
	"github.com/labstack/echo/v4"
)

//JwtConfig JwtConfig
var JwtConfig middleware.JWTConfig

func init() {
	JwtConfig = middleware.JWTConfig{
		Claims:     &dto.JwtCustomClaims{},
		SigningKey: []byte(viper.GetString("jwtSign")),
		ErrorHandler: func(err error) error {
            return echo.NewHTTPError(http.StatusUnauthorized, "Unauthorized: "+err.Error())
        },
	}
}
