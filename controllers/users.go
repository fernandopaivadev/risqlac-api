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
			Message: "Error parsing body params",
			Error:   err,
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

func DeleteUser(context *fiber.Ctx) error {
	var query DeleteQuery
	err := context.QueryParser(&query)

	if err != nil {
		context.Status(fiber.StatusBadRequest).JSON(ErrorResponse{
			Message: "Error parsing query params",
			Error:   err,
		})
		return err
	}

	slice := database.Users
	var indexToDelete uint64

	for index, user := range slice {
		if user.Id == query.Id {
			indexToDelete = uint64(index)
		}
	}

	copy(slice[indexToDelete:], slice[indexToDelete+1:])
	slice = slice[:len(slice)-1]

	database.Users = slice

	return context.SendStatus(fiber.StatusOK)
}
