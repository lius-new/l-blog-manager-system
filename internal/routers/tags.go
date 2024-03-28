package routers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/lius-new/liusnew-blog-backend-server/internal/models"
)

func RegisterTagsHanlder(app *fiber.App) {
	api := app.Group("/api/tag")

	api.Get("/view", viewHandler)
}

func viewHandler(ctx *fiber.Ctx) error {
	tags := models.ViewTags()
	return ctx.JSON(fiber.Map{"data": tags, "status": true})
}
