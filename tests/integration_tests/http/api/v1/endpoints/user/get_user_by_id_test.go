package user_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	authEndpoint "go-port-and-adapter/apps/http/api/v1/endpoints/auth"
	authHandler "go-port-and-adapter/domains/entities/auth"
	handlerDto "go-port-and-adapter/ports/domain/dto"
	repositoryMock "go-port-and-adapter/tests/unit_tests/mocks/adapters/repository/mysql"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestGetUserById(t *testing.T) {
	e := echo.New()

	repository := repositoryMock.NewUserRepository()
	repository.On("GetUserByUsername", mock.Anything)
	authHandler := authHandler.New(repository)
	authApi := authEndpoint.New(authHandler)

	claims := &handlerDto.JwtCustomClaims{
		ID: 1,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 72).Unix(),
		},
	}

	h := middleware.JWTWithConfig(middleware.JWTConfig{
		Claims:     claims,
		SigningKey: []byte("jwtSign"),
		ErrorHandler: func(err error) error {
			return echo.NewHTTPError(http.StatusUnauthorized, "Unauthorized: "+err.Error())
		},
	})(authApi.SignIn)

	reqBody, _ := json.Marshal(authEndpoint.SignInRequest{
		Username: "admin@admin.com",
		Password: "123456",
	})

	t.Run("Expect get user by id Unauthorized", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodPost, "/api/v1/auth/signin", strings.NewReader(string(reqBody)))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		req.Header.Set(echo.HeaderAuthorization, "Bearer 223434")
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		assert.Error(t, h(c))
	})

	t.Run("Expect get user by id success", func(t *testing.T) {
		encToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
		signToken, _ := encToken.SignedString([]byte("jwtSign"))

		req := httptest.NewRequest(http.MethodPost, "/api/v1/auth/signin", strings.NewReader(string(reqBody)))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		req.Header.Set(echo.HeaderAuthorization, "Bearer "+signToken)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		if assert.NoError(t, h(c)) {
			assert.Equal(t, http.StatusOK, rec.Code)
		}
	})
}
