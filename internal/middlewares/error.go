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
	case strings.Contains(err.Error(), "cannot unmarshal object"):
		return ctx.Status(fiber.StatusBadRequest).SendString("params error")
	case strings.Contains(err.Error(), "unexpected end of JSON input"): // 无法解析请求参数(未携带参数)
		return ctx.Status(fiber.StatusBadRequest).SendString("params error")
	case strings.Contains(err.Error(), " no documents in result"): // 资源不存在
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{"status": false, "message": "资源不存在"})
	case strings.Contains(err.Error(), "Method Not Allowed"): //请求不被允许
		return ctx.Status(fiber.StatusBadRequest).SendString("Method Not Allowed")
	case strings.Contains(err.Error(), errors.ErrorUnauthorized):
		return ctx.Status(fiber.StatusUnauthorized).SendString(errors.ErrorUnauthorized)
	case strings.Contains(err.Error(), errors.ErrorBlocked):
		return ctx.Status(fiber.StatusTooManyRequests).SendString(errors.ErrorBlocked)
	default:
		return ctx.Status(fiber.StatusInternalServerError).SendString("Internal Server Error: " + err.Error())
	}

}
