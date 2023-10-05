package user

import (
	"encoding/json"
	"net/http"

	"github.com/labstack/echo/v4"
)

// @Summary      Create User
// @Description  Create User
// @Tags         user
// @Param Authorization header string true "Authorization header"
// @Param id   path int true "User ID"
// @Success 200 {object} user.ReadDataResponseUser
// @Router /user/{id} [get]
func (endpoint *UserEndpoint) ReadData(c echo.Context) error {
	id := c.Param("id")

	user, err := endpoint.userHandler.ReadData(id)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, err)
	}

	b, _ := json.Marshal(user)
	readDataResponse := ReadDataResponseUser{
		ID:    user.ID,
		Username:  user.Username,
		Password: user.Password,
	}
	json.Unmarshal(b, &readDataResponse)

	return c.JSON(http.StatusOK, readDataResponse)
}

type ReadDataResponseUser struct {
	ID    int    `json:"id"`
	Username  string `json:"username"`
	Password string `json:"password"`
}
