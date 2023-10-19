package user_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	userEndpoint "go-port-and-adapter/apps/http/api/v1/endpoints/user"
	userHandler "go-port-and-adapter/domains/entities/user"
	repositoryMock "go-port-and-adapter/tests/unit_tests/mocks/adapters/repository/mysql"
	storageMock "go-port-and-adapter/tests/unit_tests/mocks/adapters/storage"

	"github.com/golang-jwt/jwt/v5"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"go-port-and-adapter/apps/http/api/v1/middleware"
)

func TestListUser(t *testing.T) {
	e := echo.New()

	repository := repositoryMock.NewUserRepository()
	repository.On("ListUser", mock.Anything)
	repository.On("GetUserById", mock.Anything)
	storage := storageMock.NewStorageManager()

	userHandler := userHandler.New(repository, storage)
	userApi := userEndpoint.New(userHandler)

	claims := &middleware.JwtCustomClaims{
		1,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 72)),
		},
	}

	h := echojwt.WithConfig(echojwt.Config{
		NewClaimsFunc: func(c echo.Context) jwt.Claims {
			return new(middleware.JwtCustomClaims)
		},
		SigningKey: []byte("jwtSign"),
		ErrorHandler: func(c echo.Context, err error) error {
			if err != nil {
				return c.JSON(http.StatusUnauthorized, "Unauthorized: "+err.Error())
			}
			return nil
		},
	})(userApi.ListUser)

	t.Run("Expect list user Unautorized", func(t *testing.T) {
		reqBody, _ := json.Marshal(userEndpoint.ListUserRequest{
			Filter:   make(map[string]string),
			OrderBy:  make(map[string]string),
			PageSize: 1,
			Page:     1,
		})

		req := httptest.NewRequest(http.MethodPost, "/api/v1/user", strings.NewReader(string(reqBody)))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		req.Header.Set(echo.HeaderAuthorization, "Bearer 223434")
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		if assert.NoError(t, h(c)) {
			assert.Equal(t, http.StatusUnauthorized, rec.Code)
		}
	})

	t.Run("Expect list success", func(t *testing.T) {
		encToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
		signToken, _ := encToken.SignedString([]byte("jwtSign"))

		reqBody, _ := json.Marshal(userEndpoint.ListUserRequest{
			Filter:   make(map[string]string),
			OrderBy:  make(map[string]string),
			PageSize: 1,
			Page:     1,
		})

		req := httptest.NewRequest(http.MethodPost, "/api/v1/user", strings.NewReader(string(reqBody)))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		req.Header.Set(echo.HeaderAuthorization, "Bearer "+signToken)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		if assert.NoError(t, h(c)) {
			assert.Equal(t, http.StatusOK, rec.Code)
		}
	})
}
