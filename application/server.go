package application

import (
	"errors"
	"net/http"
	"risqlac-api/environment"

	"github.com/labstack/echo/v4"
	echoMiddleware "github.com/labstack/echo/v4/middleware"
)

type server struct {
	Instance *echo.Echo
}

var Server server

func (server *server) Setup() {
	server.Instance = echo.New()
	server.Instance.Use(echoMiddleware.Recover())
	server.Instance.Use(echoMiddleware.Logger())
	server.Instance.Use(echoMiddleware.RequestID())
	server.Instance.Use(echoMiddleware.GzipWithConfig(echoMiddleware.GzipConfig{
		Level: 5,
	}))
	server.Instance.Use(echoMiddleware.CORSWithConfig(echoMiddleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{
			http.MethodGet,
			http.MethodPut,
			http.MethodPatch,
			http.MethodPost,
			http.MethodDelete,
			http.MethodHead,
			http.MethodOptions,
		},
		AllowHeaders: []string{
			echo.HeaderOrigin,
			echo.HeaderContentType,
			echo.HeaderAccept,
			echo.HeaderAuthorization,
		},
	}))
	server.Instance.Use(echoMiddleware.SecureWithConfig(echoMiddleware.SecureConfig{
		XSSProtection:      "1; mode=block",
		XFrameOptions:      "deny",
		ContentTypeNosniff: "nosniff",
	}))
}

func (server *server) Start() error {
	serverPort := ":" + environment.Variables.ServerPort

	err := server.Instance.Start(serverPort)

	if err != nil {
		return errors.New("error starting server: " + err.Error())
	}

	return nil
}
