package routes

import (
	"go-port-and-adapter/apps/http/api/v1/middleware"

	authEndpoint "go-port-and-adapter/apps/http/api/v1/endpoints/auth"
	userEndpoint "go-port-and-adapter/apps/http/api/v1/endpoints/user"

	authHandler "go-port-and-adapter/domains/entities/auth"
	userHandler "go-port-and-adapter/domains/entities/user"

	repository "go-port-and-adapter/adapters/repository/mysql"
	storage "go-port-and-adapter/adapters/storage"

	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
)

func API(e *echo.Echo) {
	// Init Repository Adapter
	db := repository.MysqlConnect("mysql_iam")
	userRepository := repository.NewUserRepository(db)

	// Init Storage Adapter
	storageManager := storage.NewStorageManager()

	// Init Handler
	userHandler := userHandler.New(userRepository, storageManager)
	authHandler := authHandler.New(userRepository)

	// Init Endpoint
	authEndpoint := authEndpoint.New(authHandler)
	auth := e.Group("/api/v1/auth")
	auth.POST("/signin", authEndpoint.SignIn)

	userEndpoint := userEndpoint.New(userHandler)
	user := e.Group("/api/v1/user")
	user.POST("", userEndpoint.CreateUser, echojwt.WithConfig(middleware.JwtConfig))
	user.GET("/:id", userEndpoint.GetUserById, echojwt.WithConfig(middleware.JwtConfig))
	user.PUT("/:id", userEndpoint.UpdateUser, echojwt.WithConfig(middleware.JwtConfig))
	user.DELETE("/:id", userEndpoint.DeleteUser, echojwt.WithConfig(middleware.JwtConfig))
	user.POST("/list", userEndpoint.ListUser, echojwt.WithConfig(middleware.JwtConfig))
}
