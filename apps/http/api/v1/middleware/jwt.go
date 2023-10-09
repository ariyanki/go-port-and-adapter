package middleware

import (
	"net/http"

	"go-port-and-adapter/systems/config"

	"github.com/labstack/echo/v4/middleware"
	"github.com/spf13/viper"
	"github.com/labstack/echo/v4"
)

//JwtConfig JwtConfig
var JwtConfig middleware.JWTConfig

func init() {
	config.Load()
	JwtConfig = middleware.JWTConfig{
		SigningKey: []byte(viper.GetString("jwtSign")),
		ErrorHandler: func(err error) error {
            return echo.NewHTTPError(http.StatusUnauthorized, "Unauthorized: "+err.Error())
        },
	}
}
