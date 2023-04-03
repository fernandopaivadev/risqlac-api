package controllers

import (
	"risqlac-api/application/models"
	"risqlac-api/application/services"
	"strconv"

	"github.com/labstack/echo/v4"
)

type userController struct{}

var User userController

func (*userController) Login(context echo.Context) error {
	email := context.QueryParam("email")
	password := context.QueryParam("password")

	if email == "" || password == "" {
		return context.JSON(400, echo.Map{
			"message": "bad request",
			"error":   "email and password are required",
		})
	}

	token, err := services.User.GenerateLoginToken(email, password)

	if err != nil {
		return context.JSON(401, echo.Map{
			"message": "error validating credentials",
			"error":   err.Error(),
		})
	}

	return context.JSON(200, echo.Map{
		"token": token,
	})
}

func (*userController) RequestPasswordChange(context echo.Context) error {
	email := context.QueryParam("email")

	if email == "" {
		return context.JSON(400, echo.Map{
			"message": "bad request",
			"error":   "Email is required",
		})
	}

	user, err := services.User.GetByEmail(email)

	if err != nil {
		return context.JSON(404, echo.Map{
			"message": "error retrieving user",
			"error":   err.Error(),
		})
	}

	token, err := services.User.GeneratePasswordChangeToken(user.Email)

	if err != nil {
		return context.JSON(500, echo.Map{
			"message": "error generating password change token",
			"error":   err.Error(),
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
		return context.JSON(500, echo.Map{
			"message": "error sending email",
			"error":   err.Error(),
		})
	}

	return context.JSON(200, echo.Map{
		"message": "password change token sent",
	})
}

func (*userController) ChangePassword(context echo.Context) error {
	headers := context.Request().Header
	tokenUserId, err := strconv.ParseUint(headers["Userid"][0], 10, 64)

	if err != nil {
		return context.JSON(500, echo.Map{
			"message": "error parsing token user id",
			"error":   err.Error(),
		})
	}

	password := context.QueryParam("password")

	if password == "" {
		return context.JSON(400, echo.Map{
			"message": "bad request",
			"error":   "password is required",
		})
	}

	err = services.User.ChangePassword(tokenUserId, password)

	if err != nil {
		return context.JSON(500, echo.Map{
			"message": "error changing user password",
			"error":   err.Error(),
		})
	}

	return context.JSON(200, echo.Map{
		"message": "user password changed",
	})
}

func (*userController) Create(context echo.Context) error {
	headers := context.Request().Header
	isAdmin := headers["Isadmin"][0] == "true"

	var user models.User
	err := context.Bind(&user)

	if err != nil {
		return context.JSON(400, echo.Map{
			"message": "bad request",
			"error":   err.Error(),
		})
	}

	if !isAdmin {
		user.IsAdmin = false
	}

	err = services.Utils.ValidateStruct(user)

	if err != nil {
		return context.JSON(400, echo.Map{
			"message": "bad request",
			"error":   err.Error(),
		})
	}

	err = services.User.Create(user)

	if err != nil {
		return context.JSON(500, echo.Map{
			"message": "error creating user",
			"error":   err.Error(),
		})
	}

	return context.JSON(200, echo.Map{
		"message": "user created",
	})
}

func (*userController) Update(context echo.Context) error {
	headers := context.Request().Header
	isAdmin := headers["Isadmin"][0] == "true"
	tokenUserId, err := strconv.ParseUint(headers["Userid"][0], 10, 64)

	if err != nil {
		return context.JSON(500, echo.Map{
			"message": "error parsing token user id",
			"error":   err.Error(),
		})
	}

	var user models.User
	err = context.Bind(&user)

	if err != nil {
		return context.JSON(500, echo.Map{
			"message": "error parsing body",
			"error":   err.Error(),
		})
	}

	if !(isAdmin || tokenUserId == user.Id) {
		return context.JSON(403, echo.Map{
			"message": "not allowed for no admin users",
		})
	}

	user.Password = "..." // needs a not empty value to pass validation
	err = services.Utils.ValidateStruct(user)

	if err != nil {
		return context.JSON(400, echo.Map{
			"message": "bad request",
			"error":   err.Error(),
		})
	}

	err = services.User.Update(user)

	if err != nil {
		return context.JSON(500, echo.Map{
			"message": "error updating user",
			"error":   err.Error(),
		})
	}

	return context.JSON(200, echo.Map{
		"message": "user updated",
	})
}

func (*userController) List(context echo.Context) error {
	headers := context.Request().Header
	isAdmin := headers["Isadmin"][0] == "true"
	tokenUserId, err := strconv.ParseUint(headers["Userid"][0], 10, 64)

	if err != nil {
		return context.JSON(500, echo.Map{
			"message": "error parsing token user id",
			"error":   err.Error(),
		})
	}

	if !isAdmin {
		user, err := services.User.GetById(tokenUserId)

		if err != nil {
			return context.JSON(500, echo.Map{
				"message": "error retrieving user",
				"error":   err.Error(),
			})
		}

		return context.JSON(200, echo.Map{
			"users": []models.User{user},
		})
	}

	userId, _ := strconv.ParseUint(context.QueryParam("id"), 10, 64)

	if userId != 0 {
		user, err := services.User.GetById(userId)

		if err != nil {
			return context.JSON(500, echo.Map{
				"message": "error retrieving user",
				"error":   err.Error(),
			})
		}

		return context.JSON(200, echo.Map{
			"users": []models.User{user},
		})
	}

	users, err := services.User.List()

	if err != nil {
		return context.JSON(500, echo.Map{
			"message": "error retrieving users",
			"error":   err.Error(),
		})
	}

	return context.JSON(200, echo.Map{
		"users": users,
	})
}

func (*userController) Delete(context echo.Context) error {
	headers := context.Request().Header
	isAdmin := headers["Isadmin"][0] == "true"
	tokenUserId, err := strconv.ParseUint(headers["Userid"][0], 10, 64)

	if err != nil {
		return context.JSON(500, echo.Map{
			"message": "error parsing token user id",
			"error":   err.Error(),
		})
	}

	userId, err := strconv.ParseUint(context.QueryParam("id"), 10, 64)

	if err != nil {
		return context.JSON(400, echo.Map{
			"message": "bad request",
			"error":   "id must be a number",
		})
	}

	if !(isAdmin || tokenUserId == userId) {
		return context.JSON(403, echo.Map{
			"message": "not allowed for no admin users",
		})
	}

	if err != nil {
		return context.JSON(400, echo.Map{
			"message": "bad request",
			"error":   err.Error(),
		})
	}

	err = services.User.Delete(userId)

	if err != nil {
		return context.JSON(500, echo.Map{
			"message": "error deleting user",
			"error":   err.Error(),
		})
	}

	return context.JSON(200, echo.Map{
		"message": "user deleted",
	})
}
