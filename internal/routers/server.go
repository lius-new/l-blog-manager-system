package routers

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/csrf"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/lius-new/liusnew-blog-backend-server/internal/middlewares"
)

func Server() {
	app := fiber.New(fiber.Config{ErrorHandler: middlewares.FiberConfigErrorHandler})

	app.Use(recover.New())
	app.Use(csrf.New(csrf.Config{ErrorHandler: middlewares.CrosErrrHandler}))

	app.Use(middlewares.CorsMiddlware)
	app.Use(middlewares.BaseLoggerMiddleware)
	app.Use(middlewares.AuthMiddleware)

	app.Get("/time", func(c *fiber.Ctx) error {
		return c.SendString(time.Now().String())
	})

	// 注册路由
	RegisterUserHanlder(app)
	RegisterArticlesHanlder(app)
	RegisterTagsHanlder(app)

	app.Listen(":8080")
}
