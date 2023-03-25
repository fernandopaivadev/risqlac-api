package controllers

import (
	"fmt"
	"risqlac-api/models"
	"risqlac-api/services"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type userController struct{}

var User userController

func (*userController) Login(context *fiber.Ctx) error {
	var query userAuthRequest
	err := context.QueryParser(&query)

	if err != nil {
		return context.Status(fiber.StatusInternalServerError).JSON(errorResponse{
			Message: "Error parsing query params",
			Error:   err.Error(),
		})
	}

	err = services.Utils.ValidateStruct(query)

	if err != nil {
		return context.Status(fiber.StatusBadRequest).JSON(errorResponse{
			Message: "Bad request",
			Error:   err.Error(),
		})
	}

	token, err := services.User.GenerateLoginToken(query.Email, query.Password)

	if err != nil {
		return context.Status(fiber.StatusInternalServerError).JSON(errorResponse{
			Message: "Error generating token",
			Error:   err.Error(),
		})
	}

	return context.Status(fiber.StatusOK).JSON(userAuthResponse{
		Token: token,
	})
}

func (*userController) RequestPasswordChange(context *fiber.Ctx) error {
	var query requestPasswordChangeRequest
	err := context.QueryParser(&query)

	if err != nil {
		return context.Status(fiber.StatusInternalServerError).JSON(errorResponse{
			Message: "Error parsing query params",
			Error:   err.Error(),
		})
	}

	err = services.Utils.ValidateStruct(query)

	if err != nil {
		return context.Status(fiber.StatusBadRequest).JSON(errorResponse{
			Message: "Bad request",
			Error:   err.Error(),
		})
	}

	user, err := services.User.GetByEmail(query.Email)

	if err != nil {
		return context.Status(fiber.StatusInternalServerError).JSON(errorResponse{
			Message: "Error retrieving user",
			Error:   err.Error(),
		})
	}

	token, err := services.User.GeneratePasswordChangeToken(user.Email)

	if err != nil {
		return context.Status(fiber.StatusInternalServerError).JSON(errorResponse{
			Message: "Error generating password change token",
			Error:   err.Error(),
		})
	}

	err = services.Utils.SendEmail(
		user.Name,
		user.Email,
		"RECUPERAÇÃO DE SENHA",
		"TOKEN DE RECUPERAÇÃO: "+token,
		"",
	)

	if err != nil {
		return context.Status(fiber.StatusInternalServerError).JSON(errorResponse{
			Message: "Error sending email",
			Error:   err.Error(),
		})
	}

	return context.Status(fiber.StatusAccepted).JSON(messageResponse{
		Message: "Password recovery email sent",
	})
}

func (*userController) ChangePassword(context *fiber.Ctx) error {
	requestHeaders := context.GetReqHeaders()
	tokenUserId, err := strconv.ParseUint(requestHeaders["Userid"], 10, 64)

	if err != nil {
		return context.Status(fiber.StatusInternalServerError).JSON(errorResponse{
			Message: "Error parsing token user id",
			Error:   err.Error(),
		})
	}

	var query changePasswordRequest
	err = context.QueryParser(&query)

	if err != nil {
		return context.Status(fiber.StatusInternalServerError).JSON(errorResponse{
			Message: "Error parsing query params",
			Error:   err.Error(),
		})
	}

	err = services.Utils.ValidateStruct(query)

	if err != nil {
		return context.Status(fiber.StatusBadRequest).JSON(errorResponse{
			Message: "Bad request",
			Error:   err.Error(),
		})
	}

	err = services.User.ChangePassword(tokenUserId, query.Password)

	if err != nil {
		return context.Status(fiber.StatusInternalServerError).JSON(errorResponse{
			Message: "Error changing user password",
			Error:   err.Error(),
		})
	}

	return context.Status(fiber.StatusOK).JSON(messageResponse{
		Message: "User password changed",
	})
}

func (*userController) Create(context *fiber.Ctx) error {
	requestHeaders := context.GetReqHeaders()
	isAdmin := requestHeaders["Isadmin"] == "true"

	var user models.User
	err := context.BodyParser(&user)

	if err != nil {
		return context.Status(fiber.StatusBadRequest).JSON(errorResponse{
			Message: "Error parsing body params",
			Error:   err.Error(),
		})
	}

	if !isAdmin {
		user.IsAdmin = false
	}

	err = services.Utils.ValidateStruct(user)

	if err != nil {
		return context.Status(fiber.StatusBadRequest).JSON(errorResponse{
			Message: "Bad request",
			Error:   err.Error(),
		})
	}

	err = services.User.Create(user)

	if err != nil {
		return context.Status(fiber.StatusInternalServerError).JSON(errorResponse{
			Message: "Error creating user",
			Error:   err.Error(),
		})
	}

	return context.Status(fiber.StatusCreated).JSON(messageResponse{
		Message: "User created",
	})
}

func (*userController) Update(context *fiber.Ctx) error {
	requestHeaders := context.GetReqHeaders()
	isAdmin := requestHeaders["Isadmin"] == "true"
	tokenUserId, err := strconv.ParseUint(requestHeaders["Userid"], 10, 64)

	if err != nil {
		return context.Status(fiber.StatusInternalServerError).JSON(errorResponse{
			Message: "Error parsing token user id",
			Error:   err.Error(),
		})
	}

	var user models.User
	err = context.BodyParser(&user)

	if err != nil {
		return context.Status(fiber.StatusInternalServerError).JSON(errorResponse{
			Message: "Error parsing body params",
			Error:   err.Error(),
		})
	}

	if !(isAdmin || tokenUserId == user.Id) {
		return context.Status(fiber.StatusForbidden).JSON(messageResponse{
			Message: "Not allowed for no admin users",
		})
	}

	user.Password = "..." // needs a not empty value to pass validation
	err = services.Utils.ValidateStruct(user)

	if err != nil {
		return context.Status(fiber.StatusBadRequest).JSON(errorResponse{
			Message: "Bad request",
			Error:   err.Error(),
		})
	}

	err = services.User.Update(user)

	if err != nil {
		return context.Status(fiber.StatusInternalServerError).JSON(errorResponse{
			Message: "Error updating user",
			Error:   err.Error(),
		})
	}

	return context.Status(fiber.StatusOK).JSON(messageResponse{
		Message: "User updated",
	})
}

func (*userController) List(context *fiber.Ctx) error {
	requestHeaders := context.GetReqHeaders()
	isAdmin := requestHeaders["Isadmin"] == "true"
	tokenUserId, err := strconv.ParseUint(requestHeaders["Userid"], 10, 64)

	fmt.Println(requestHeaders)

	if err != nil {
		return context.Status(fiber.StatusInternalServerError).JSON(errorResponse{
			Message: "Error parsing token user id",
			Error:   err.Error(),
		})
	}

	var query byIdRequest
	err = context.QueryParser(&query)

	if err != nil {
		return context.Status(fiber.StatusInternalServerError).JSON(errorResponse{
			Message: "Error parsing query params",
			Error:   err.Error(),
		})
	}

	var users []models.User

	if isAdmin {
		if query.Id != 0 {
			user, err := services.User.GetById(query.Id)

			if err != nil {
				return context.Status(fiber.StatusInternalServerError).JSON(errorResponse{
					Message: "Error retrieving user",
					Error:   err.Error(),
				})
			}

			users = append(users, user)
		} else {
			users, err = services.User.List()

			if err != nil {
				return context.Status(fiber.StatusInternalServerError).JSON(errorResponse{
					Message: "Error retrieving users",
					Error:   err.Error(),
				})
			}
		}
	} else {
		user, err := services.User.GetById(tokenUserId)

		if err != nil {
			return context.Status(fiber.StatusInternalServerError).JSON(errorResponse{
				Message: "Error retrieving user",
				Error:   err.Error(),
			})
		}

		users = append(users, user)
	}

	if err != nil {
		return context.Status(fiber.StatusInternalServerError).JSON(errorResponse{
			Message: "Error retrieving users",
			Error:   err.Error(),
		})
	}

	return context.Status(fiber.StatusOK).JSON(listUsersResponse{
		Users: users,
	})
}

func (*userController) Delete(context *fiber.Ctx) error {
	requestHeaders := context.GetReqHeaders()
	isAdmin := requestHeaders["Isadmin"] == "true"
	tokenUserId, err := strconv.ParseUint(requestHeaders["Userid"], 10, 64)

	if err != nil {
		return context.Status(fiber.StatusInternalServerError).JSON(errorResponse{
			Message: "Error parsing token user id",
			Error:   err.Error(),
		})
	}

	var query byIdRequest
	err = context.QueryParser(&query)

	if err != nil {
		return context.Status(fiber.StatusInternalServerError).JSON(errorResponse{
			Message: "Error parsing query params",
			Error:   err.Error(),
		})
	}

	if !(isAdmin || tokenUserId == query.Id) {
		return context.Status(fiber.StatusForbidden).JSON(messageResponse{
			Message: "Not allowed for no admin users",
		})
	}

	err = services.Utils.ValidateStruct(query)

	if err != nil {
		return context.Status(fiber.StatusBadRequest).JSON(errorResponse{
			Message: "Bad request",
			Error:   err.Error(),
		})
	}

	err = services.User.Delete(query.Id)

	if err != nil {
		return context.Status(fiber.StatusInternalServerError).JSON(errorResponse{
			Message: "Error deleting user",
			Error:   err.Error(),
		})
	}

	return context.Status(fiber.StatusOK).JSON(messageResponse{
		Message: "User deleted",
	})
}
