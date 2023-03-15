package controllers

import (
	"risqlac-api/services"
	"risqlac-api/types"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func UserLogin(context *fiber.Ctx) error {
	var query types.UserAuthRequest
	err := context.QueryParser(&query)

	if err != nil {
		return context.Status(fiber.StatusInternalServerError).JSON(types.ErrorResponse{
			Message: "Error parsing query params",
			Error:   err.Error(),
		})
	}

	err = services.ValidateStruct(query)

	if err != nil {
		return context.Status(fiber.StatusBadRequest).JSON(types.ErrorResponse{
			Message: "Bad request",
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

func RequestPasswordChange(context *fiber.Ctx) error {
	var query types.RequestPasswordChangeRequest
	err := context.QueryParser(&query)

	if err != nil {
		return context.Status(fiber.StatusInternalServerError).JSON(types.ErrorResponse{
			Message: "Error parsing query params",
			Error:   err.Error(),
		})
	}

	err = services.ValidateStruct(query)

	if err != nil {
		return context.Status(fiber.StatusBadRequest).JSON(types.ErrorResponse{
			Message: "Bad request",
			Error:   err.Error(),
		})
	}

	user, err := services.GetUserByEmail(query.Email)

	if err != nil {
		return context.Status(fiber.StatusInternalServerError).JSON(types.ErrorResponse{
			Message: "Error retrieving user",
			Error:   err.Error(),
		})
	}

	token, err := services.GeneratePasswordChangeToken(user.Email)

	if err != nil {
		return context.Status(fiber.StatusInternalServerError).JSON(types.ErrorResponse{
			Message: "Error generating password change token",
			Error:   err.Error(),
		})
	}

	go services.SendEmail(
		user.Name,
		user.Email,
		"RECUPERAÇÃO DE SENHA",
		"TOKEN DE RECUPERAÇÃO: "+token,
		"",
	)

	return context.Status(fiber.StatusAccepted).JSON(types.SuccessResponse{
		Message: "Password recovery email sent",
	})
}

func ChangePassword(context *fiber.Ctx) error {
	requestHeaders := context.GetReqHeaders()
	tokenUserId, err := strconv.ParseUint(requestHeaders["User_id"], 10, 64)

	if err != nil {
		return context.Status(fiber.StatusInternalServerError).JSON(types.ErrorResponse{
			Message: "Error parsing token user id",
			Error:   err.Error(),
		})
	}

	var query types.ChangePasswordRequest
	err = context.QueryParser(&query)

	if err != nil {
		return context.Status(fiber.StatusInternalServerError).JSON(types.ErrorResponse{
			Message: "Error parsing query params",
			Error:   err.Error(),
		})
	}

	err = services.ValidateStruct(query)

	if err != nil {
		return context.Status(fiber.StatusBadRequest).JSON(types.ErrorResponse{
			Message: "Bad request",
			Error:   err.Error(),
		})
	}

	err = services.ChangeUserPassword(tokenUserId, query.Password)

	if err != nil {
		return context.Status(fiber.StatusInternalServerError).JSON(types.ErrorResponse{
			Message: "Error changing user password",
			Error:   err.Error(),
		})
	}

	return context.Status(fiber.StatusOK).JSON(types.SuccessResponse{
		Message: "User password changed",
	})
}

func CreateUser(context *fiber.Ctx) error {
	requestHeaders := context.GetReqHeaders()
	isAdmin := requestHeaders["Is_admin"] == "true"

	var user types.User
	err := context.BodyParser(&user)

	if err != nil {
		return context.Status(fiber.StatusBadRequest).JSON(types.ErrorResponse{
			Message: "Error parsing body params",
			Error:   err.Error(),
		})
	}

	if !isAdmin {
		user.IsAdmin = false
	}

	err = services.ValidateStruct(user)

	if err != nil {
		return context.Status(fiber.StatusBadRequest).JSON(types.ErrorResponse{
			Message: "Bad request",
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
	isAdmin := requestHeaders["Is_admin"] == "true"
	tokenUserId, err := strconv.ParseUint(requestHeaders["User_id"], 10, 64)

	if err != nil {
		return context.Status(fiber.StatusInternalServerError).JSON(types.ErrorResponse{
			Message: "Error parsing token user id",
			Error:   err.Error(),
		})
	}

	var user types.User
	err = context.BodyParser(&user)

	if err != nil {
		return context.Status(fiber.StatusInternalServerError).JSON(types.ErrorResponse{
			Message: "Error parsing body params",
			Error:   err.Error(),
		})
	}

	if !(isAdmin || tokenUserId == user.Id) {
		return context.Status(fiber.StatusForbidden).JSON(types.MessageResponse{
			Message: "Not allowed for no admin users",
		})
	}

	user.Password = "..." // needs a not empty value to pass validation
	err = services.ValidateStruct(user)

	if err != nil {
		return context.Status(fiber.StatusBadRequest).JSON(types.ErrorResponse{
			Message: "Bad request",
			Error:   err.Error(),
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
	tokenUserId, err := strconv.ParseUint(requestHeaders["User_id"], 10, 64)

	if err != nil {
		return context.Status(fiber.StatusInternalServerError).JSON(types.ErrorResponse{
			Message: "Error parsing token user id",
			Error:   err.Error(),
		})
	}

	var query types.ByIdRequest
	err = context.QueryParser(&query)

	if err != nil {
		return context.Status(fiber.StatusInternalServerError).JSON(types.ErrorResponse{
			Message: "Error parsing query params",
			Error:   err.Error(),
		})
	}

	var users []types.User

	if isAdmin {
		if query.Id != 0 {
			user, err := services.GetUserById(query.Id)

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
		user, err := services.GetUserById(tokenUserId)

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
	tokenUserId, err := strconv.ParseUint(requestHeaders["User_id"], 10, 64)

	if err != nil {
		return context.Status(fiber.StatusInternalServerError).JSON(types.ErrorResponse{
			Message: "Error parsing token user id",
			Error:   err.Error(),
		})
	}

	var query types.ByIdRequest
	err = context.QueryParser(&query)

	if err != nil {
		return context.Status(fiber.StatusInternalServerError).JSON(types.ErrorResponse{
			Message: "Error parsing query params",
			Error:   err.Error(),
		})
	}

	if !(isAdmin || tokenUserId == query.Id) {
		return context.Status(fiber.StatusForbidden).JSON(types.MessageResponse{
			Message: "Not allowed for no admin users",
		})
	}

	err = services.ValidateStruct(query)

	if err != nil {
		return context.Status(fiber.StatusBadRequest).JSON(types.ErrorResponse{
			Message: "Bad request",
			Error:   err.Error(),
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
