package main

import (
	"risqlac-api/controllers"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/requestid"
)

type Response struct {
	Message string `json:"message"`
}

func main() {
	app := fiber.New()
	app.Use(logger.New())
	app.Use(requestid.New())

	app.Get("/info", func(context *fiber.Ctx) error {
		return context.Status(fiber.StatusOK).JSON(Response{
			Message: "RisQLAC API",
		})
	})

	user := app.Group("/user")
	product := app.Group("/product")

	user.Post("/create", controllers.CreateUser)
	user.Get("/list", controllers.ListUsers)
	user.Delete("/delete", controllers.DeleteUser)

	product.Post("/create", controllers.CreateProduct)
	product.Get("/list", controllers.ListProducts)
	product.Delete("/delete", controllers.DeleteProduct)

	app.Listen(":3000")
}
