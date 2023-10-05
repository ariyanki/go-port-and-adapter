package config

import (
	"github.com/labstack/echo/v4/middleware"
	"github.com/spf13/viper"
	"go-port-and-adapter/ports/system/dto"
)

//JwtConfig JwtConfig
var JwtConfig middleware.JWTConfig

func init() {
	JwtConfig = middleware.JWTConfig{
		Claims:     &dto.JwtCustomClaims{},
		SigningKey: []byte(viper.GetString("jwtSign")),
	}
}
