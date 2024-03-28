package routers

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/lius-new/liusnew-blog-backend-server/internal/models"
)

func RegisterArticlesHanlder(app *fiber.App) {
	api := app.Group("/api/articles")

	api.Post("/create", createHander)
	api.Post("/modify", modifyHander)
	api.Post("/delete", deleteHander)
	api.Post("/view", viewHander)
}

func createHander(ctx *fiber.Ctx) error {
	type article struct {
		Title   string   `json:"title" bind:"required"`
		Content string   `json:"content" bind:"required"`
		Tags    []string `json:"tags" bind:"required"`
		Covers  []string `json:"covers"`
	}
	a := new(article)
	if err := ctx.BodyParser(a); err != nil {
		log.Println(err)
		return err
	}

	tags := models.SaveTags(a.Tags)

	article_, err := models.CreateArticles(a.Title, a.Content, tags, a.Covers)
	if err != nil && err.Error() == "article exist" {
		return ctx.JSON(fiber.Map{"message": "article exist"})
	}

	article_.Tags = a.Tags

	return ctx.JSON(fiber.Map{"data": article_, "status": true})
}

func modifyHander(ctx *fiber.Ctx) error {
	type article struct {
		Id      string   `json:"id"`
		Title   string   `json:"title"`
		Content string   `json:"content"`
		Tags    []string `json:"tags"`
		Covers  []string `json:"covers"`
		Status  bool     `json:"status"`
	}

	a := new(article)
	if err := ctx.BodyParser(a); err != nil {
		log.Println(err)
		return err
	}

	tags := models.SaveTags(a.Tags)

	article_, err := models.ModifyArticles(a.Id, a.Title, a.Content, tags, a.Covers, a.Status)
	if err != nil && err.Error() == "article not found" {
		return ctx.JSON(fiber.Map{"message": "article not found"})
	}

	article_.Tags = a.Tags

	return ctx.JSON(fiber.Map{"data": article_})
}

func deleteHander(ctx *fiber.Ctx) error {
	type article struct {
		Id   string   `json:"id"`
		Tags []string `json:"tags"`
	}

	a := new(article)
	if err := ctx.BodyParser(a); err != nil {
		return err
	}

	models.DeleteTags(a.Tags)
	models.DeleteArticles(a.Id)

	return ctx.JSON(fiber.Map{"message": "delete success"})
}
func viewHander(ctx *fiber.Ctx) error {
	type article struct {
		PageSize int64 `json:"page_size"`
		PageNum  int64 `json:"page_num"`
	}

	a := new(article)
	if err := ctx.BodyParser(a); err != nil {
		return err
	}

	articles, count := models.ViewArticles(a.PageSize, a.PageNum)

	return ctx.JSON(fiber.Map{"data": articles, "total": count, "status": true})
}
