package user_test

import (
	"net/http"
	"net/http/httptest"
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

func TestGetUserById(t *testing.T) {
	e := echo.New()

	repository := repositoryMock.NewUserRepository()
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
	})(userApi.GetUserById)

	t.Run("Expect get user by id Unautorized", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/api/v1/user", nil)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		req.Header.Set(echo.HeaderAuthorization, "Bearer 223434")
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetParamNames("id")
		c.SetParamValues("1")

		if assert.NoError(t, h(c)) {
			assert.Equal(t, http.StatusUnauthorized, rec.Code)
		}
	})

	t.Run("Expect get user by id success", func(t *testing.T) {
		encToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
		signToken, _ := encToken.SignedString([]byte("jwtSign"))

		req := httptest.NewRequest(http.MethodGet, "/api/v1/user", nil)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		req.Header.Set(echo.HeaderAuthorization, "Bearer "+signToken)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetParamNames("id")
		c.SetParamValues("1")

		if assert.NoError(t, h(c)) {
			assert.Equal(t, http.StatusOK, rec.Code)
		}
	})

	t.Run("Expect get user by id not found", func(t *testing.T) {
		encToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
		signToken, _ := encToken.SignedString([]byte("jwtSign"))

		req := httptest.NewRequest(http.MethodGet, "/api/v1/user", nil)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		req.Header.Set(echo.HeaderAuthorization, "Bearer "+signToken)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetParamNames("id")
		c.SetParamValues("123")

		if assert.NoError(t, h(c)) {
			assert.Equal(t, http.StatusUnprocessableEntity, rec.Code)
		}
	})
}
