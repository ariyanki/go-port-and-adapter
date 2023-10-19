package user

import (
	"net/http"
	"time"

	handlerDto "go-port-and-adapter/ports/domain/dto"
	"go-port-and-adapter/systems/automap"
	"go-port-and-adapter/systems/validator"

	"github.com/labstack/echo/v4"
)

// @Summary      List User
// @Description  List User
// @Tags         user
// @Param Authorization header string true "Authorization header"
// @Param request body user.ListUserRequest true "List User Request Body"
// @Success 200
// @Router /user/list [post]
func (endpoint *UserEndpoint) ListUser(c echo.Context) error {
	// read request param
	listRequest := new(ListUserRequest)
	if err := c.Bind(listRequest); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, err.Error())
	}
	// validate request param
	if err := validator.GetValidator().Struct(listRequest); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	// map request param to request handler
	var requestHandler handlerDto.ListRequestDto
	if err := automap.AutoMap(listRequest, &requestHandler); err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	// call handler
	var responseHandler []handlerDto.UserDto
	if err := endpoint.userHandler.ListUser(requestHandler, &responseHandler); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, err.Error())
	}

	// map response handler to response http
	var response []handlerDto.UserDto
	if err := automap.AutoMap(responseHandler, &response); err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, response)
}

type ListUserRequest struct {
	Filter   map[string]string `json:"filter"`
	OrderBy  map[string]string `json:"order_by"`
	PageSize int               `json:"page_size"`
	Page     int               `json:"page"`
}

type ListUserResponse struct {
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
