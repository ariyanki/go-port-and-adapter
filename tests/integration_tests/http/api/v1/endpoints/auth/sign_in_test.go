package auth_test

import (
    "testing"
	"net/http"
	"net/http/httptest"
	"strings"
	"encoding/json"

	authHandler "go-port-and-adapter/domains/entities/auth"
	authEndpoint "go-port-and-adapter/apps/http/api/v1/endpoints/auth"
	repositoryMock "go-port-and-adapter/tests/unit_tests/mocks/adapters/repository/mysql"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestSignIn(t *testing.T) {
	t.Run("Expect sign in success", func(t *testing.T) {
		e := echo.New()

		repository := repositoryMock.NewUserRepository()
		repository.On("GetUserByUsername", mock.Anything)

		authHandler := authHandler.New(repository)
		authApi := authEndpoint.New(authHandler)

		reqBody, _ := json.Marshal(authEndpoint.SignInRequest{
			Username: "admin@admin.com",
			Password: "123456",
		})

		req := httptest.NewRequest(http.MethodPost, "/api/v1/auth/signin", strings.NewReader(string(reqBody)))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		
		if assert.NoError(t, authApi.SignIn(c)) {
			assert.Equal(t, http.StatusOK, rec.Code)
		}
	})
}