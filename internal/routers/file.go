package routers

import (
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/lius-new/liusnew-blog-backend-server/internal/utils"
)

func RegisterFileHanlder(app *fiber.App) {
	api := app.Group("/api/file")
	api.Get("/:hash", viewFileHandle)
	api.Get("/hello", func(c *fiber.Ctx) error {
		return c.SendString("hello")
	})
}

func viewFileHandle(ctx *fiber.Ctx) error {
	param := struct {
		Hash string `params:"hash"`
	}{}

	if err := ctx.ParamsParser(&param); err != nil {
		return ctx.SendStatus(fiber.ErrBadRequest.Code)
	}

	var filename string
	if filename, _ = utils.FileExist(param.Hash, os.Getenv("COVER_PATH")); len(filename) == 0 {
		return ctx.SendStatus(fiber.ErrNotFound.Code)
	}

	return ctx.SendFile(os.Getenv("COVER_PATH") + filename)
}
