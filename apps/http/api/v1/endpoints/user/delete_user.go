package user

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

// @Summary      Delete User
// @Description  Delete User
// @Tags         user
// @Param Authorization header string true "Authorization header"
// @Param id   path int true "User ID"
// @Success 200
// @Router /user/{id} [delete]
func (endpoint *UserEndpoint) DeleteUser(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, err.Error())
	}
	if err := endpoint.userHandler.DeleteUser(id); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, err.Error())
	}

	return c.JSON(http.StatusOK, "")
}
