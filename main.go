package main

import "github.com/gofiber/fiber/v2"

type Response struct {
	message string
}

func main() {
	server := fiber.New()

	server.Get("/", func(context *fiber.Ctx) error {
		data := Response{
			message: "AE GAROTO",
		}
		return context.Status(fiber.StatusOK).JSON(data)
	})

	server.Listen(":3000")
}
