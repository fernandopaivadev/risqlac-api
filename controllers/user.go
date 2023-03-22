package controllers

import (
	"risqlac-api/services"
	"risqlac-api/types"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type userController struct{}

var User userController

func (*userController) Login(context *fiber.Ctx) error {
	var query types.UserAuthRequest
	err := context.QueryParser(&query)

	if err != nil {
		return context.Status(fiber.StatusInternalServerError).JSON(types.ErrorResponse{
			Message: "Error parsing query params",
			Error:   err.Error(),
		})
	}

	err = services.Utils.ValidateStruct(query)

	if err != nil {
		return context.Status(fiber.StatusBadRequest).JSON(types.ErrorResponse{
			Message: "Bad request",
			Error:   err.Error(),
		})
	}

	token, err := services.User.GenerateLoginToken(query.Email, query.Password)

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

func (*userController) RequestPasswordChange(context *fiber.Ctx) error {
	var query types.RequestPasswordChangeRequest
	err := context.QueryParser(&query)

	if err != nil {
		return context.Status(fiber.StatusInternalServerError).JSON(types.ErrorResponse{
			Message: "Error parsing query params",
			Error:   err.Error(),
		})
	}

	err = services.Utils.ValidateStruct(query)

	if err != nil {
		return context.Status(fiber.StatusBadRequest).JSON(types.ErrorResponse{
			Message: "Bad request",
			Error:   err.Error(),
		})
	}

	user, err := services.User.GetByEmail(query.Email)

	if err != nil {
		return context.Status(fiber.StatusInternalServerError).JSON(types.ErrorResponse{
			Message: "Error retrieving user",
			Error:   err.Error(),
		})
	}

	token, err := services.User.GeneratePasswordChangeToken(user.Email)

	if err != nil {
		return context.Status(fiber.StatusInternalServerError).JSON(types.ErrorResponse{
			Message: "Error generating password change token",
			Error:   err.Error(),
		})
	}

	go func() {
		_ = services.Utils.SendEmail(
			user.Name,
			user.Email,
			"RECUPERAÇÃO DE SENHA",
			"TOKEN DE RECUPERAÇÃO: "+token,
			"",
		)
	}()

	return context.Status(fiber.StatusAccepted).JSON(types.SuccessResponse{
		Message: "Password recovery email sent",
	})
}

func (*userController) ChangePassword(context *fiber.Ctx) error {
	requestHeaders := context.GetReqHeaders()
	tokenUserId, err := strconv.ParseUint(requestHeaders["UserId"], 10, 64)

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

	err = services.Utils.ValidateStruct(query)

	if err != nil {
		return context.Status(fiber.StatusBadRequest).JSON(types.ErrorResponse{
			Message: "Bad request",
			Error:   err.Error(),
		})
	}

	err = services.User.ChangePassword(tokenUserId, query.Password)

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

func (*userController) Create(context *fiber.Ctx) error {
	requestHeaders := context.GetReqHeaders()
	isAdmin := requestHeaders["IsAdmin"] == "true"

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

	err = services.Utils.ValidateStruct(user)

	if err != nil {
		return context.Status(fiber.StatusBadRequest).JSON(types.ErrorResponse{
			Message: "Bad request",
			Error:   err.Error(),
		})
	}

	err = services.User.Create(user)

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

func (*userController) Update(context *fiber.Ctx) error {
	requestHeaders := context.GetReqHeaders()
	isAdmin := requestHeaders["IsAdmin"] == "true"
	tokenUserId, err := strconv.ParseUint(requestHeaders["UserId"], 10, 64)

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
	err = services.Utils.ValidateStruct(user)

	if err != nil {
		return context.Status(fiber.StatusBadRequest).JSON(types.ErrorResponse{
			Message: "Bad request",
			Error:   err.Error(),
		})
	}

	err = services.User.Update(user)

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

func (*userController) List(context *fiber.Ctx) error {
	requestHeaders := context.GetReqHeaders()
	isAdmin := requestHeaders["IsAdmin"] == "true"
	tokenUserId, err := strconv.ParseUint(requestHeaders["UserId"], 10, 64)

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
			user, err := services.User.GetById(query.Id)

			if err != nil {
				return context.Status(fiber.StatusInternalServerError).JSON(types.ErrorResponse{
					Message: "Error retrieving user",
					Error:   err.Error(),
				})
			}

			users = append(users, user)
		} else {
			users, err = services.User.List()

			if err != nil {
				return context.Status(fiber.StatusInternalServerError).JSON(types.ErrorResponse{
					Message: "Error retrieving users",
					Error:   err.Error(),
				})
			}
		}
	} else {
		user, err := services.User.GetById(tokenUserId)

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

func (*userController) Delete(context *fiber.Ctx) error {
	requestHeaders := context.GetReqHeaders()
	isAdmin := requestHeaders["IsAdmin"] == "true"
	tokenUserId, err := strconv.ParseUint(requestHeaders["UserId"], 10, 64)

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

	err = services.Utils.ValidateStruct(query)

	if err != nil {
		return context.Status(fiber.StatusBadRequest).JSON(types.ErrorResponse{
			Message: "Bad request",
			Error:   err.Error(),
		})
	}

	err = services.User.Delete(query.Id)

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
