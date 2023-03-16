package main

import (
	"risqlac-api/environment"
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

type Server struct {
	App *fiber.App
}

var server Server

func (service *Server) Setup() {
	service.App = fiber.New(fiber.Config{
		JSONEncoder: json.Marshal,
		JSONDecoder: json.Unmarshal,
	})

	service.App.Use(recover.New())
	service.App.Use(logger.New())
	service.App.Use(requestid.New())
	service.App.Use(compress.New(compress.Config{
		Level: compress.LevelBestSpeed,
	}))

	service.App.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowMethods: "GET,POST,HEAD,PUT,DELETE,PATCH,OPTIONS",
	}))

	service.App.Get("/info", func(context *fiber.Ctx) error {
		return context.Status(fiber.StatusOK).SendString("RisQLAC API v2.3.1")
	})

	service.App.Get("/metrics", monitor.New(monitor.Config{
		Title:   "RisQLAC API Metrics",
		Refresh: time.Second * 5,
	}))
}

func (service *Server) Start() {
	serverPort := ":" + environment.Variables.ServerPort

	err := service.App.Listen(serverPort)

	if err != nil {
		panic("Error starting Server: " + err.Error())
	}
}
