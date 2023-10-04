package user

import (
	"encoding/json"
	"net/http"

	"github.com/labstack/echo/v4"
)

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
