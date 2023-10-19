package user

import (
	"net/http"
	"strconv"
	"time"

	"go-port-and-adapter/apps/http/api/v1/middleware"
	handlerDto "go-port-and-adapter/ports/domain/dto"
	"go-port-and-adapter/systems/automap"
	"go-port-and-adapter/systems/validator"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

// @Summary      Update User
// @Description  Update User
// @Tags         user
// @Param Authorization header string true "Authorization header"
// @Param id   path int true "User ID"
// @Param request body user.UpdateUserRequest true "Update User Request Body"
// @Success 200
// @Router /user/{id} [put]
func (endpoint *UserEndpoint) UpdateUser(c echo.Context) error {
	// Get JWT claims
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*middleware.JwtCustomClaims)

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, err.Error())
	}

	request := new(UpdateUserRequest)

	if err := c.Bind(request); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, err.Error())
	}

	if err := validator.GetValidator().Struct(request); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	var requestHandler handlerDto.UpdateUserDto
	if err := automap.AutoMap(request, &requestHandler); err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	requestHandler.ID = id
	requestHandler.UpdatedBy = claims.ID

	if err := endpoint.userHandler.UpdateUser(requestHandler); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, err.Error())
	}

	return c.JSON(http.StatusOK, "")
}

type UpdateUserRequest struct {
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
