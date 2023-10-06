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
// @Success 200 {object} user.GetUserByIdResponseUser
// @Router /user/{id} [get]
func (endpoint *UserEndpoint) GetUserById(c echo.Context) error {
	id := c.Param("id")

	user, err := endpoint.userHandler.GetUserById(id)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, err)
	}

	b, _ := json.Marshal(user)
	getUserByIdResponse := GetUserByIdResponseUser{
		ID:    user.ID,
		Username:  user.Username,
		Password: user.Password,
	}
	json.Unmarshal(b, &getUserByIdResponse)

	return c.JSON(http.StatusOK, getUserByIdResponse)
}

type GetUserByIdResponseUser struct {
	ID    int    `json:"id"`
	Username  string `json:"username"`
	Password string `json:"password"`
}
