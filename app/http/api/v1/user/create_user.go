package user


import (
	"net/http"

	"go-port-and-adapter/helpers/validator"
	domainDto "go-port-and-adapter/ports/domain/dto"

	"github.com/labstack/echo/v4"
)

func (controller *UserController) CreateUser(c echo.Context) error {
	createRequest := new(CreateUserRequest)

	if err := c.Bind(createRequest); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, err)
	}

	if err := validator.GetValidator().Struct(createRequest); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	user := domainDto.CreateUserDto{
		Username:  createRequest.Username,
		Password: createRequest.Password,
	}

	if err := controller.userDomain.CreateUser(user); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, err)
	}

	return c.JSON(http.StatusOK, "")
}


type CreateUserRequest struct {
	Username  string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}