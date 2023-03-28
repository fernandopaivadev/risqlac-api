package application

import (
	"errors"
	"risqlac-api/infra"
	"time"

	"github.com/goccy/go-json"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/monitor"
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

	server.App.Get("/info", func(context *fiber.Ctx) error {
		return context.Status(fiber.StatusOK).SendString("RisQLAC API v2.4.20")
	})

	server.App.Get("/metrics", monitor.New(monitor.Config{
		Title:   "RisQLAC API Metrics",
		Refresh: time.Second * 5,
	}))
}

func (server *server) Start() error {
	serverPort := ":" + infra.Environment.Variables.ServerPort

	err := server.App.Listen(serverPort)

	if err != nil {
		return errors.New("error starting server: " + err.Error())
	}

	return nil
}
