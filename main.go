package main

import (
	"risqlac-api/controllers"
	"risqlac-api/database"
	"risqlac-api/environment"
	"risqlac-api/middleware"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/requestid"
)

func main() {
	environment.Load()
	database.Connect()

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
	userRoutes.Get(
		"/auth",
		controllers.UserAuth,
	)
	userRoutes.Post(
		"/create",
		controllers.CreateUser,
	)
	userRoutes.Put(
		"/update",
		middleware.ValidateToken,
		controllers.UpdateUser,
	)
	userRoutes.Get(
		"/list",
		middleware.ValidateToken,
		middleware.VerifyAdmin,
		controllers.ListUsers,
	)
	userRoutes.Delete(
		"/delete",
		middleware.ValidateToken,
		middleware.VerifyAdmin,
		controllers.DeleteUser,
	)

	productRoutes := app.Group("/product")
	productRoutes.Post(
		"/create",
		middleware.ValidateToken,
		controllers.CreateProduct,
	)
	productRoutes.Put(
		"/update",
		middleware.ValidateToken,
		controllers.UpdateProduct,
	)
	productRoutes.Get(
		"/list",
		middleware.ValidateToken,
		controllers.ListProducts,
	)
	productRoutes.Delete(
		"/delete",
		middleware.ValidateToken,
		controllers.DeleteProduct,
	)

	err := app.Listen(":3000")

	if err != nil {
		panic("Error starting server: " + err.Error())
	}
}
