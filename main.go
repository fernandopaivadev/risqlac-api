package main

import (
	"risqlac-api/controllers"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/requestid"
)

func main() {
	app := fiber.New()

	app.Use(logger.New())
	app.Use(requestid.New())
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "Origin, Content-Type, Accept",
	}))

	app.Get("/info", func(context *fiber.Ctx) error {
		return context.Status(fiber.StatusOK).SendString("RisQLAC API v1.0")
	})

	userRoutes := app.Group("/user")
	productRoutes := app.Group("/product")

	userRoutes.Post("/create", controllers.CreateUser)
	userRoutes.Get("/list", controllers.ListUsers)
	userRoutes.Delete("/delete", controllers.DeleteUser)

	productRoutes.Post("/create", controllers.CreateProduct)
	productRoutes.Get("/list", controllers.ListProducts)
	productRoutes.Delete("/delete", controllers.DeleteProduct)

	app.Listen(":3000")
}
