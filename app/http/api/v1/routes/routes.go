package routes

import (
	"go-port-and-adapter/helpers/config"

	authEndpoint "go-port-and-adapter/app/http/api/v1/endpoints/auth"
	userEndpoint "go-port-and-adapter/app/http/api/v1/endpoints/user"
	
	authHandler "go-port-and-adapter/domain/entities/auth"
	userHandler "go-port-and-adapter/domain/entities/user"

	repository "go-port-and-adapter/adapters/repository/mysql"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func API(e *echo.Echo) {

	// Init Repository
	db := config.DB
	userRepository := repository.NewUserRepository(db)

	// Init Handler
	userHandler := userHandler.New(userRepository)
	authHandler := authHandler.New(userRepository)

	// Init Endpoint
	authEndpoint := authEndpoint.New(authHandler)
	auth := e.Group("/api/v1/auth/signin")
	auth.POST("", authEndpoint.SignIn)
	
	userEndpoint := userEndpoint.New(userHandler)
	user := e.Group("/api/v1/user")
	user.POST("", userEndpoint.CreateUser, middleware.JWTWithConfig(config.JwtConfig))
	user.GET("/:id", userEndpoint.ReadData, middleware.JWTWithConfig(config.JwtConfig))
}
