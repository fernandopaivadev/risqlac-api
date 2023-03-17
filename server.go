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

func (server *Server) Setup() {
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
		return context.Status(fiber.StatusOK).SendString("RisQLAC API v2.4.0")
	})

	server.App.Get("/metrics", monitor.New(monitor.Config{
		Title:   "RisQLAC API Metrics",
		Refresh: time.Second * 5,
	}))
}

func (server *Server) Start() {
	serverPort := ":" + environment.Variables.ServerPort

	err := server.App.Listen(serverPort)

	if err != nil {
		panic("Error starting Server: " + err.Error())
	}
}
