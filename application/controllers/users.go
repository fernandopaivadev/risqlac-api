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
			"message": "validation error",
			"error":   "email and password are required",
		})
	}

	token, err := services.User.GenerateSessionToken(email, password)

	if err != nil {
		return context.JSON(500, echo.Map{
			"message": "error generating session token",
			"error":   err.Error(),
		})
	}

	return context.JSON(200, echo.Map{
		"token": token,
	})
}

func (*userController) ListSessions(context echo.Context) error {
	headers := context.Request().Header
	tokenUserId, err := strconv.ParseUint(headers["Userid"][0], 10, 64)

	if err != nil {
		return context.JSON(500, echo.Map{
			"message": "error parsing token user id",
			"error":   err.Error(),
		})
	}

	sessions, err := services.Session.GetByUserId(tokenUserId)

	if err != nil {
		return context.JSON(500, echo.Map{
			"message": "error retrieving sessions",
			"error":   err.Error(),
		})
	}

	return context.JSON(200, echo.Map{
		"sessions": sessions,
	})
}

func (*userController) Logout(context echo.Context) error {
	headers := context.Request().Header
	token := headers["Authorization"][0]

	err := services.Session.DeleteByToken(token)

	if err != nil {
		return context.JSON(500, echo.Map{
			"message": "error logging out",
			"error":   err.Error(),
		})
	}

	return context.JSON(200, echo.Map{
		"message": "user successfully logged out",
	})
}

func (*userController) CompleteLogout(context echo.Context) error {
	headers := context.Request().Header
	tokenUserId, err := strconv.ParseUint(headers["Userid"][0], 10, 64)

	if err != nil {
		return context.JSON(500, echo.Map{
			"message": "error parsing token user id",
			"error":   err.Error(),
		})
	}

	err = services.Session.DeleteByUserId(tokenUserId)

	if err != nil {
		return context.JSON(500, echo.Map{
			"message": "error logging out",
			"error":   err.Error(),
		})
	}

	return context.JSON(200, echo.Map{
		"message": "user successfully logged out of all sessions",
	})
}

func (*userController) Create(context echo.Context) error {
	headers := context.Request().Header
	isAdmin := headers["Isadmin"][0] == "1"

	var user models.User
	err := context.Bind(&user)

	if err != nil {
		return context.JSON(500, echo.Map{
			"message": "error parsing body",
			"error":   err.Error(),
		})
	}

	if !isAdmin {
		user.IsAdmin = 0
	}

	err = services.Utils.ValidateStruct(user)

	if err != nil {
		return context.JSON(400, echo.Map{
			"message": "validation error",
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

	return context.JSON(201, echo.Map{
		"message": "user created",
	})
}

func (*userController) Update(context echo.Context) error {
	headers := context.Request().Header
	isAdmin := headers["Isadmin"][0] == "1"
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
			"message": "not allowed for not admin users",
		})
	}

	user.Password = "..." // needs a not empty value to pass validation
	err = services.Utils.ValidateStruct(user)

	if err != nil {
		return context.JSON(400, echo.Map{
			"message": "validation error",
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
	isAdmin := headers["Isadmin"][0] == "1"
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
	isAdmin := headers["Isadmin"][0] == "1"
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
			"message": "validation error",
			"error":   err.Error(),
		})
	}

	if !(isAdmin || tokenUserId == userId) {
		return context.JSON(403, echo.Map{
			"message": "not allowed for not admin users",
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
