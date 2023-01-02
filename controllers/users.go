package controllers

import (
	"risqlac-api/models"
	"risqlac-api/services"
	"risqlac-api/types"
	"strconv"

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
	requestHeaders := context.GetReqHeaders()
	isAdmin := requestHeaders["Is_admin"] == "true"

	var user models.User
	err := context.BodyParser(&user)

	if err != nil {
		return context.Status(fiber.StatusBadRequest).JSON(types.ErrorResponse{
			Message: "Error parsing body params",
			Error:   err.Error(),
		})
	}

	if !isAdmin {
		user.Is_admin = false
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
	isAdmin := requestHeaders["Is_admin"] == "true"
	loggedUserId, err := strconv.ParseUint(requestHeaders["User_id"], 10, 64)

	if err != nil {
		return context.Status(fiber.StatusInternalServerError).JSON(types.ErrorResponse{
			Message: "Error parsing token user id",
			Error:   err.Error(),
		})
	}

	var user models.User
	err = context.BodyParser(&user)

	if err != nil {
		return context.Status(fiber.StatusInternalServerError).JSON(types.ErrorResponse{
			Message: "Error parsing body params",
			Error:   err.Error(),
		})
	}

	if !(isAdmin || loggedUserId == user.Id) {
		return context.Status(fiber.StatusForbidden).JSON(types.MessageResponse{
			Message: "Not allowed for no admin users",
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
	isAdmin := requestHeaders["Is_admin"] == "true"
	loggedUserId, err := strconv.ParseUint(requestHeaders["User_id"], 10, 64)

	if err != nil {
		return context.Status(fiber.StatusInternalServerError).JSON(types.ErrorResponse{
			Message: "Error parsing token user id",
			Error:   err.Error(),
		})
	}

	var query types.QueryById
	err = context.QueryParser(&query)

	if err != nil {
		return context.Status(fiber.StatusInternalServerError).JSON(types.ErrorResponse{
			Message: "Error parsing query params",
			Error:   err.Error(),
		})
	}

	var users []models.User

	if isAdmin {
		if query.Id != 0 {
			user, err := services.GetUser(query.Id)

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
	} else {
		user, err := services.GetUser(loggedUserId)

		if err != nil {
			return context.Status(fiber.StatusInternalServerError).JSON(types.ErrorResponse{
				Message: "Error retrieving user",
				Error:   err.Error(),
			})
		}

		users = append(users, user)
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
	requestHeaders := context.GetReqHeaders()
	isAdmin := requestHeaders["Is_admin"] == "true"
	loggedUserId, err := strconv.ParseUint(requestHeaders["User_id"], 10, 64)

	if err != nil {
		return context.Status(fiber.StatusInternalServerError).JSON(types.ErrorResponse{
			Message: "Error parsing token user id",
			Error:   err.Error(),
		})
	}

	var query types.QueryById
	err = context.QueryParser(&query)

	if err != nil {
		return context.Status(fiber.StatusInternalServerError).JSON(types.ErrorResponse{
			Message: "Error parsing query params",
			Error:   err.Error(),
		})
	}

	if !(isAdmin || loggedUserId == query.Id) {
		return context.Status(fiber.StatusForbidden).JSON(types.MessageResponse{
			Message: "Not allowed for no admin users",
		})
	}

	err = services.DeleteUser(query.Id)

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
