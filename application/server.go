package application

import (
	"errors"
	"risqlac-api/environment"

	"github.com/goccy/go-json"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/fiber/v2/middleware/requestid"
)

type server struct {
	App *fiber.App
}

var Server server

func (server *server) Setup() {
	server.App = fiber.New(fiber.Config{
		JSONEncoder: json.Marshal,
		JSONDecoder: json.Unmarshal,
	})

	server.App.Use(recover.New())
	server.App.Use(logger.New())
	server.App.Use(requestid.New())
	server.App.Use(compress.New(compress.Config{
		Level: compress.LevelBestSpeed,
	}))

	server.App.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowMethods: "GET,POST,HEAD,PUT,DELETE,PATCH,OPTIONS",
	}))
}

func (server *server) Start() error {
	serverPort := ":" + environment.Variables.ServerPort

	err := server.App.Listen(serverPort)

	if err != nil {
		return errors.New("error starting server: " + err.Error())
	}

	return nil
}
