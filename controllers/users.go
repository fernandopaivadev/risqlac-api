package controllers

import (
	"risqlac-api/models"
	"risqlac-api/services"

	"github.com/gofiber/fiber/v2"
)

func UserAuth(context *fiber.Ctx) error {
	var query UserAuthQuery
	err := context.QueryParser(&query)

	if err != nil {
		return context.Status(fiber.StatusInternalServerError).JSON(ErrorResponse{
			Message: "Error parsing query params",
			Error:   err.Error(),
		})
	}

	token, err := services.GenerateUserToken(query.Email, query.Password)

	if err != nil {
		return context.Status(fiber.StatusInternalServerError).JSON(ErrorResponse{
			Message: "Error generating token",
			Error:   err.Error(),
		})
	}

	return context.Status(fiber.StatusOK).JSON(UserAuthResponse{
		Token: token,
	})
}

func CreateUser(context *fiber.Ctx) error {
	var user models.User
	err := context.BodyParser(&user)

	if err != nil {
		return context.Status(fiber.StatusBadRequest).JSON(ErrorResponse{
			Message: "Error parsing body params",
			Error:   err.Error(),
		})
	}

	err = services.CreateUser(user)

	if err != nil {
		return context.Status(fiber.StatusInternalServerError).JSON(ErrorResponse{
			Message: "Error creating user",
			Error:   err.Error(),
		})
	}

	return context.Status(fiber.StatusCreated).JSON(SuccessResponse{
		Message: "User created",
	})
}

func UpdateUser(context *fiber.Ctx) error {
	var user models.User
	err := context.BodyParser(&user)

	if err != nil {
		return context.Status(fiber.StatusInternalServerError).JSON(ErrorResponse{
			Message: "Error parsing body params",
			Error:   err.Error(),
		})
	}

	err = services.UpdateUser(user)

	if err != nil {
		return context.Status(fiber.StatusInternalServerError).JSON(ErrorResponse{
			Message: "Error updating user",
			Error:   err.Error(),
		})
	}

	return context.Status(fiber.StatusOK).JSON(SuccessResponse{
		Message: "User updated",
	})
}

func ListUsers(context *fiber.Ctx) error {
	var query QueryById
	err := context.QueryParser(&query)

	if err != nil {
		return context.Status(fiber.StatusInternalServerError).JSON(ErrorResponse{
			Message: "Error parsing query params",
			Error:   err.Error(),
		})
	}

	users, err := services.ListUsers(query.Id)

	if err != nil {
		return context.Status(fiber.StatusInternalServerError).JSON(ErrorResponse{
			Message: "Error retrieving users",
			Error:   err.Error(),
		})
	}

	return context.Status(fiber.StatusOK).JSON(ListUsersResponse{
		Users: users,
	})
}

func DeleteUser(context *fiber.Ctx) error {
	var query QueryById
	err := context.QueryParser(&query)

	if err != nil {
		return context.Status(fiber.StatusInternalServerError).JSON(ErrorResponse{
			Message: "Error parsing query params",
			Error:   err.Error(),
		})
	}

	err = services.DeleteUser(query.Id)

	if err != nil {
		return context.Status(fiber.StatusInternalServerError).JSON(ErrorResponse{
			Message: "Error deleting user",
			Error:   err.Error(),
		})
	}

	return context.Status(fiber.StatusOK).JSON(SuccessResponse{
		Message: "User deleted",
	})
}
