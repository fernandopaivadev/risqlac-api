package controllers

import (
	"risqlac-api/database"
	"risqlac-api/models"

	"github.com/gofiber/fiber/v2"
)

func CreateUser(context *fiber.Ctx) error {
	var user models.User
	err := context.BodyParser(&user)

	if err != nil {
		context.Status(fiber.StatusBadRequest).JSON(ErrorResponse{
			Error: err,
		})
		return err
	}

	database.Users = append(database.Users, user)

	return context.Status(fiber.StatusCreated).JSON(CreatedUserResponse{
		CreatedUser: user,
	})
}

func ListUsers(context *fiber.Ctx) error {
	return context.Status(fiber.StatusCreated).JSON(ListUsersResponse{
		Users: database.Users,
	})
}
