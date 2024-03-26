package internal

import (
	"github.com/gofiber/fiber/v2"
	"github.com/lius-new/liusnew-blog-backend-server/internal/middlewares"
)

func Server() {
	app := fiber.New()

	app.Use(middlewares.BaseLoggerMiddleware)

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("this is test")
	})

	app.Listen(":8080")
}
