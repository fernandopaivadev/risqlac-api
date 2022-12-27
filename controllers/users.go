package controllers

import (
	"fmt"
	"risqlac-api/models"
	"risqlac-api/services"
	"risqlac-api/types"

	"github.com/gofiber/fiber/v2"
)

func UserLogin(context *fiber.Ctx) error {
	var query types.UserAuthQuery
	err := context.QueryParser(&query)

	if err != nil {
		return context.Status(fiber.StatusInternalServerError).JSON(types.ErrorResponse{
			Message: "Error parsing query params",
			Error:   err.Error(),
		})
	}

	token, err := services.GenerateUserToken(query.Email, query.Password)

	if err != nil {
		return context.Status(fiber.StatusInternalServerError).JSON(types.ErrorResponse{
			Message: "Error generating token",
			Error:   err.Error(),
		})
	}

	return context.Status(fiber.StatusOK).JSON(types.UserAuthResponse{
		Token: token,
	})
}

func CreateUser(context *fiber.Ctx) error {
	var user models.User
	err := context.BodyParser(&user)

	if err != nil {
		return context.Status(fiber.StatusBadRequest).JSON(types.ErrorResponse{
			Message: "Error parsing body params",
			Error:   err.Error(),
		})
	}

	err = services.CreateUser(user)

	if err != nil {
		return context.Status(fiber.StatusInternalServerError).JSON(types.ErrorResponse{
			Message: "Error creating user",
			Error:   err.Error(),
		})
	}

	return context.Status(fiber.StatusCreated).JSON(types.SuccessResponse{
		Message: "User created",
	})
}

func UpdateUser(context *fiber.Ctx) error {
	requestHeaders := context.GetReqHeaders()
	isAdmin := requestHeaders["is_admin"] == "true"
	loggedUserId := requestHeaders["user_id"]

	var user models.User
	err := context.BodyParser(&user)

	if err != nil {
		return context.Status(fiber.StatusInternalServerError).JSON(types.ErrorResponse{
			Message: "Error parsing body params",
			Error:   err.Error(),
		})
	}

	if !(isAdmin || loggedUserId == fmt.Sprint(user.Id)) {
		return context.Status(fiber.StatusForbidden).JSON(types.MessageResponse{
			Message: "User is not admin",
		})
	}

	err = services.UpdateUser(user)

	if err != nil {
		return context.Status(fiber.StatusInternalServerError).JSON(types.ErrorResponse{
			Message: "Error updating user",
			Error:   err.Error(),
		})
	}

	return context.Status(fiber.StatusOK).JSON(types.SuccessResponse{
		Message: "User updated",
	})
}

func ListUsers(context *fiber.Ctx) error {
	requestHeaders := context.GetReqHeaders()
	isAdmin := requestHeaders["is_admin"] == "true"
	loggedUserId := requestHeaders["user_id"]

	var query types.ListUsersQuery
	err := context.QueryParser(&query)

	if err != nil {
		return context.Status(fiber.StatusInternalServerError).JSON(types.ErrorResponse{
			Message: "Error parsing query params",
			Error:   err.Error(),
		})
	}

	if !(isAdmin || loggedUserId == fmt.Sprint(query.UserId)) {
		return context.Status(fiber.StatusForbidden).JSON(types.MessageResponse{
			Message: "User is not admin",
		})
	}

	var users []models.User

	if query.UserId != 0 {
		user, err := services.GetUser(query.UserId)

		if err != nil {
			return context.Status(fiber.StatusInternalServerError).JSON(types.ErrorResponse{
				Message: "Error retrieving user",
				Error:   err.Error(),
			})
		}

		users = append(users, user)
	} else {
		users, err = services.ListUsers()

		if err != nil {
			return context.Status(fiber.StatusInternalServerError).JSON(types.ErrorResponse{
				Message: "Error retrieving users",
				Error:   err.Error(),
			})
		}
	}

	if err != nil {
		return context.Status(fiber.StatusInternalServerError).JSON(types.ErrorResponse{
			Message: "Error retrieving users",
			Error:   err.Error(),
		})
	}

	return context.Status(fiber.StatusOK).JSON(types.ListUsersResponse{
		Users: users,
	})
}

func DeleteUser(context *fiber.Ctx) error {
	var query types.DeleteUserQuery
	err := context.QueryParser(&query)

	if err != nil {
		return context.Status(fiber.StatusInternalServerError).JSON(types.ErrorResponse{
			Message: "Error parsing query params",
			Error:   err.Error(),
		})
	}

	err = services.DeleteUser(query.UserId)

	if err != nil {
		return context.Status(fiber.StatusInternalServerError).JSON(types.ErrorResponse{
			Message: "Error deleting user",
			Error:   err.Error(),
		})
	}

	return context.Status(fiber.StatusOK).JSON(types.SuccessResponse{
		Message: "User deleted",
	})
}
