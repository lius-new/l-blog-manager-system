package middlewares

import "github.com/gofiber/fiber/v2"

func CrosErrrHandler(c *fiber.Ctx, err error) error {
	return c.Next()
}
