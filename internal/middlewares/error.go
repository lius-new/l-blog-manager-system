package middlewares

import (
	"strings"

	"github.com/gofiber/fiber/v2"
)

func ErrorMiddleware(ctx *fiber.Ctx, err error) error {
	if err == nil {
		return nil
	}

	switch true {
	case strings.Contains(err.Error(), "Cannot "):
		return ctx.Status(fiber.StatusNotFound).SendString("Page Not Found")
	default:
		return ctx.Status(fiber.StatusInternalServerError).SendString("Internal Server Error!!!")
	}

}
