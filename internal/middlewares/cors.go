package middlewares

import (
	"github.com/gofiber/fiber/v2"
)

var ALLOO_RIGINS = []string{
	"http://localhost:5173",
	"http://localhost:3000",
	"http://101.42.173.48:8080",
}

func CorsMiddlware(c *fiber.Ctx) error {

	origin := c.Get("Origin")

	for _, v := range ALLOO_RIGINS {
		if v == origin {
			c.Set("Access-Control-Allow-Origin", v)
			break
		}
	}

	c.Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE")
	c.Set("Access-Control-Allow-Headers", "Content-Type, Cookie")
	c.Set("Access-Control-Allow-Credentials", "true")
	if c.Method() == "OPTIONS" {
		return c.SendStatus(fiber.StatusOK)
	}
	return c.Next()
}
