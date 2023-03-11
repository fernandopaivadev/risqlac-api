package server

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/monitor"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/fiber/v2/middleware/requestid"
)

var Instance *fiber.App

func Init() {
	Instance = fiber.New()
	//Instance = fiber.New(fiber.Config{
	//	JSONEncoder: json.Marshal,
	//	JSONDecoder: json.Unmarshal,
	//})

	Instance.Use(recover.New())
	Instance.Use(logger.New())
	Instance.Use(requestid.New())
	Instance.Use(compress.New(compress.Config{
		Level: compress.LevelBestSpeed,
	}))

	Instance.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowMethods: "GET,POST,HEAD,PUT,DELETE,PATCH,OPTIONS",
	}))

	Instance.Get("/info", func(context *fiber.Ctx) error {
		return context.Status(fiber.StatusOK).SendString("RisQLAC API v2.2.14")
	})

	Instance.Get("/metrics", monitor.New(monitor.Config{
		Title:   "RisQLAC API Metrics",
		Refresh: time.Second * 5,
	}))
}
