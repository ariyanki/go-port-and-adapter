package user

import (
	"go-port-and-adapter/apps/http/api/v1/middleware"
	handlerDto "go-port-and-adapter/ports/domain/dto"
	"time"

	"go-port-and-adapter/systems/automap"
	"go-port-and-adapter/systems/validator"
	"net/http"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

// @Summary      Create User
// @Description  Create User
// @Tags         user
// @Param Authorization header string true "Authorization header"
// @Param request body user.CreateUserRequest true "Create User Request Body"
// @Success 200
// @Router /user [post]
func (endpoint *UserEndpoint) CreateUser(c echo.Context) error {
	// Get JWT claims
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*middleware.JwtCustomClaims)

	request := new(CreateUserRequest)

	if err := c.Bind(request); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, err.Error())
	}

	if err := validator.GetValidator().Struct(request); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	var requestHandler handlerDto.CreateUserDto
	if err := automap.AutoMap(request, &requestHandler); err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	requestHandler.CreatedBy = claims.ID

	if err := endpoint.userHandler.CreateUser(requestHandler); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, err.Error())
	}

	return c.JSON(http.StatusOK, "")
}

type CreateUserRequest struct {
	Username      string    `json:"username"  validate:"required"`
	Fullname      string    `json:"fullname"  validate:"required"`
	Phonenumber   string    `json:"phonenumber"  validate:"required"`
	Email         string    `json:"email"  validate:"required"`
	PhotoFilename string    `json:"photo_filename"`
	PhotoFiledata string    `json:"photo_filedata"`
	Birthdate     time.Time `json:"birthdate"  validate:"required"`
	Gender        string    `json:"gender"  validate:"required"`
	City          string    `json:"city"  validate:"required"`
	Address       string    `json:"address"  validate:"required"`
}
