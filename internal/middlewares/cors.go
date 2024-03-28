package middlewares

import "github.com/gofiber/fiber/v2"

func CorsMiddlware(c *fiber.Ctx) error {
	c.Set("Access-Control-Allow-Origin", "http://localhost:5173")
	c.Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE")
	c.Set("Access-Control-Allow-Headers", "Content-Type, Cookie")
	c.Set("Access-Control-Allow-Credentials", "true")
	if c.Method() == "OPTIONS" {
		return c.SendStatus(fiber.StatusOK)
	}
	return c.Next()

}
