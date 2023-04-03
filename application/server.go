package application

import (
	"errors"
	"risqlac-api/environment"

	"github.com/labstack/echo/v4"
)

type server struct {
	Instance *echo.Echo
}

var Server server

func (server *server) Setup() {
	server.Instance = echo.New()
	// server.App = fiber.New(fiber.Config{
	// 	JSONEncoder: json.Marshal,
	// 	JSONDecoder: json.Unmarshal,
	// })

	// server.App.Use(recover.New())
	// server.App.Use(logger.New())
	// server.App.Use(requestid.New())
	// server.App.Use(compress.New(compress.Config{
	// 	Level: compress.LevelBestSpeed,
	// }))

	// server.App.Use(cors.New(cors.Config{
	// 	AllowOrigins: "*",
	// 	AllowMethods: "GET,POST,HEAD,PUT,DELETE,PATCH,OPTIONS",
	// }))
}

func (server *server) Start() error {
	serverPort := ":" + environment.Variables.ServerPort

	err := server.Instance.Start(serverPort)

	if err != nil {
		return errors.New("error starting server: " + err.Error())
	}

	return nil
}
