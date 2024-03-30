package routers

import (
	"os"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/csrf"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/lius-new/liusnew-blog-backend-server/internal/logger"
	"github.com/lius-new/liusnew-blog-backend-server/internal/middlewares"
)

func buildServer() *fiber.App {
	app := fiber.New(fiber.Config{
		ErrorHandler: middlewares.FiberConfigErrorHandler,
	})

	app.Use(recover.New())
	app.Use(csrf.New(csrf.Config{ErrorHandler: middlewares.CrosErrrHandler}))

	app.Use(middlewares.CorsMiddlware)
	app.Use(middlewares.BaseLoggerMiddleware)
	app.Use(middlewares.AuthMiddleware)

	return app
}

func Server1() {
	app := buildServer()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendStatus(fiber.ErrBadGateway.Code)
	})

	// 注册路由
	RegisterUserHanlder(app)
	RegisterArticlesHanlder(app)
	RegisterTagsHanlder(app)

	if err := app.Listen(strings.Join([]string{":", os.Getenv("SEVER_PORT_1")}, "")); err != nil {
		logger.Panic(err.Error())
	}
}
func Server2() {
	app := buildServer()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendStatus(fiber.ErrBadGateway.Code)
	})
	app.Get("/time", func(c *fiber.Ctx) error {
		return c.SendString(time.Now().String())
	})

	// 注册路由
	RegisterArticlesHanlder2(app)
	RegisterTagsHanlder(app)

	if err := app.Listen(strings.Join([]string{":", os.Getenv("SEVER_PORT_2")}, "")); err != nil {
		logger.Panic(err.Error())
	}
}
