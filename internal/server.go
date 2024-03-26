package internal

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/joho/godotenv"
	"github.com/lius-new/liusnew-blog-backend-server/internal/middlewares"
)

func init() {
	if err := godotenv.Load("../.env"); err != nil {
		panic("No .env file")
	}
}

func Server() {
	app := fiber.New(fiber.Config{ErrorHandler: middlewares.ErrorMiddleware})

	app.Use(recover.New())
	app.Use(middlewares.BaseLoggerMiddleware)

	app.Get("/time", func(c *fiber.Ctx) error {
		return c.SendString(time.Now().String())
	})

	app.Get("/errors", func(c *fiber.Ctx) error {
		panic("this is test")
	})

	app.Listen(":8080")
}
