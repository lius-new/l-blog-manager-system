package middlewares

import (
	"github.com/gofiber/fiber/v2"
)

func ErrorMiddleware(ctx *fiber.Ctx, err error) error {
	return ctx.Status(fiber.StatusInternalServerError).SendString("Internal Server Error!!!")
}
