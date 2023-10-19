package user

import (
	"go-port-and-adapter/systems/automap"
	"net/http"
	"strconv"
	"time"

	handlerDto "go-port-and-adapter/ports/domain/dto"

	"github.com/labstack/echo/v4"
)

// @Summary      Get User by ID
// @Description  Get User by ID
// @Tags         user
// @Param Authorization header string true "Authorization header"
// @Param id   path int true "User ID"
// @Success 200 {object} user.GetUserByIdResponse
// @Router /user/{id} [get]
func (endpoint *UserEndpoint) GetUserById(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, err.Error())
	}
	var user handlerDto.UserDto
	if err := endpoint.userHandler.GetUserById(id, &user); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, err.Error())
	}

	var response GetUserByIdResponse
	if err := automap.AutoMap(user, &response); err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, response)
}

type GetUserByIdResponse struct {
	ID                int       `json:"id"`
	Name              string    `json:"name"`
	Address           string    `json:"address"`
	Capacity          int       `json:"capacity"`
	Status            int       `json:"status"`
	CreatedByFullname string    `json:"created_by_fullname"`
	UpdatedByFullname string    `json:"updated_by_fullname"`
	CreatedAt         time.Time `json:"created_at"`
	UpdatedAt         time.Time `json:"updated_at"`
	DeletedAt         time.Time `json:"deleted_at"`
}
