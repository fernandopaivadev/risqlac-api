package middleware

import (
	"risqlac-api/types"

	"github.com/gofiber/fiber/v2"
)

func VerifyAdmin(context *fiber.Ctx) error {
	requestHeaders := context.GetReqHeaders()
	isAdmin := requestHeaders["Is_admin"] == "true"

	if !isAdmin {
		return context.Status(fiber.StatusForbidden).JSON(types.MessageResponse{
			Message: "Not allowed for no admin users",
		})
	}

	return context.Next()
}
