package middleware

import (
	"github.com/gofiber/fiber/v2"
)

func VerifyAdmin(context *fiber.Ctx) error {
	requestHeaders := context.GetReqHeaders()
	isAdmin := requestHeaders["is_admin"] == "true"

	if isAdmin {
		return context.Next()
	} else {
		return context.Status(fiber.StatusForbidden).JSON(MessageResponse{
			Message: "User is not admin",
		})
	}
}
