package middlewares

import (
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/lius-new/liusnew-blog-backend-server/internal/errors"
	"github.com/lius-new/liusnew-blog-backend-server/internal/logger"
)

func FiberConfigErrorHandler(ctx *fiber.Ctx, err error) error {
	if err == nil {
		return nil
	}
	logger.Error(err)

	switch true {
	case strings.Contains(err.Error(), "Cannot "):
		return ctx.Status(fiber.StatusNotFound).SendString("Page Not Found")
	case strings.Contains(err.Error(), errors.ErrorUnauthorized):
		return ctx.Status(fiber.StatusUnauthorized).SendString(errors.ErrorUnauthorized)
	case strings.Contains(err.Error(), errors.ErrorBlocked):
		return ctx.Status(fiber.StatusTooManyRequests).SendString(errors.ErrorBlocked)
	default:
		return ctx.Status(fiber.StatusInternalServerError).SendString("Internal Server Error: " + err.Error())
	}

}
