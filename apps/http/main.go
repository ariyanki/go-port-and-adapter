package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"time"

	"go-port-and-adapter/apps/http/api/v1/routes"
	"go-port-and-adapter/systems/config"
	"go-port-and-adapter/systems/logger"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/spf13/viper"
	echoSwagger "github.com/swaggo/echo-swagger"

	_ "go-port-and-adapter/docs"
)

// @title Swagger Go Port and Adapter API
// @version 1.0
// @description This is a Go Port and Adapter API server.

// @host localhost:7777
// @BasePath /api/v1

func main() {
	config.Load()
	e := echo.New()
	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(
		middleware.LoggerWithConfig(middleware.LoggerConfig{
			Output: logger.MiddlewareLog,
		}),
		middleware.Recover())
	e.HTTPErrorHandler = func(err error, c echo.Context) {
		logger.Error().Msgf("[ERROR] -- %s, Request Path %s", err.Error(), c.Request().URL.Path)
		e.DefaultHTTPErrorHandler(err, c)
	}
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept, echo.HeaderAuthorization},
	}))
	// Handler for hooking any request in routers registered and log it
	e.Use(middleware.BodyDumpWithConfig(middleware.BodyDumpConfig{
		Handler: logger.APILogHandler,
		Skipper: logger.APILogSkipper,
	}))
	// Handler for putting app request and response timestamp. This used for get elapsed time
	e.Use(ServiceRequestTime)

	e.GET("/swagger/*", echoSwagger.WrapHandler)

	routes.API(e)

	e.GET("/", func(c echo.Context) error {
		message := `Go Grab your Code`
		return c.String(http.StatusOK, message)
	})

	// Start server
	go func() {
		if err := e.Start(":" + viper.GetString("port")); err != nil {
			e.Logger.Info("Shutting down the server")
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server with
	// a timeout of 10 seconds.
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := e.Shutdown(ctx); err != nil {
		e.Logger.Fatal(err)
	}

}

// ServiceRequestTime middleware adds a `Server` header to the response.
func ServiceRequestTime(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		c.Request().Header.Set("X-App-RequestTime", time.Now().Format(time.RFC3339))
		return next(c)
	}
}
