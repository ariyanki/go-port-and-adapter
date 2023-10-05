package user


import (
	"net/http"

	"go-port-and-adapter/helpers/validator"
	handlerDto "go-port-and-adapter/ports/domain/dto"

	"github.com/labstack/echo/v4"
)

// @Summary      Create User
// @Description  Create User
// @Tags         user
// @Param request body user.CreateUserRequest true "Create User Request Body"
// @Success 200
// @Router /user [post]
func (endpoint *UserEndpoint) CreateUser(c echo.Context) error {
	createRequest := new(CreateUserRequest)

	if err := c.Bind(createRequest); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, err)
	}

	if err := validator.GetValidator().Struct(createRequest); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	user := handlerDto.CreateUserDto{
		Username:  createRequest.Username,
		Password: createRequest.Password,
	}

	if err := endpoint.userHandler.CreateUser(user); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, err)
	}

	return c.JSON(http.StatusOK, "")
}


type CreateUserRequest struct {
	Username  string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}