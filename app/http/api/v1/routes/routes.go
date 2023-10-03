package routes

import (
	"go-port-and-adapter/helpers/config"

	userController "go-port-and-adapter/app/http/api/v1/user"
	"go-port-and-adapter/domain"
	repository "go-port-and-adapter/adapters/repository/mysql"

	"github.com/labstack/echo/v4"
)

func API(e *echo.Echo) {

	// Init Repository
	db := config.DB
	userRepository := repository.NewUserRepository(db)

	// Init Domain
	userDomain := domain.NewUserDomain(userRepository)

	// Init Controller
	userController := userController.New(userDomain)
	user := e.Group("/api/v1/user")
	user.POST("", userController.CreateUser)
	user.GET("/:id", userController.ReadData)
}
