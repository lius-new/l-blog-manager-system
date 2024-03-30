package routers

import (
	"log"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/lius-new/liusnew-blog-backend-server/internal/models"
)

func RegisterArticlesHanlder(app *fiber.App) {
	api := app.Group("/api/articles")

	api.Post("/create", createHander)
	api.Put("/modify", modifyHander)
	api.Delete("/delete", deleteHander)
	api.Post("/views", viewsHander)
	api.Post("/view", viewHander)
}
func RegisterArticlesHanlder2(app *fiber.App) {
	api := app.Group("/api/articles")

	api.Post("/views", viewsHander)
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
		Id      string   `json:"id" bind:"required"`
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

	tags = models.ViewArticlesTags(article_.Tags)
	article_.Tags = tags

	return ctx.JSON(fiber.Map{"data": article_, "status": true})
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
func viewsHander(ctx *fiber.Ctx) error {
	type article struct {
		PageSize int64 `json:"page_size"`
		PageNum  int64 `json:"page_num"`
	}

	a := new(article)
	if err := ctx.BodyParser(a); err != nil {
		return err
	}

	articles, count := models.ViewArticles(a.PageSize, a.PageNum)

	// 获取每个文章的标签名
	for index := range articles {
		tags := models.ViewArticlesTags(articles[index].Tags)
		articles[index].Tags = tags
		if len(articles[index].Content) > 20 {
			articles[index].Content = strings.Join([]string{articles[index].Content[:20], "..."}, "")
		}
	}

	return ctx.JSON(fiber.Map{"data": articles, "total": count, "status": true})
}

func viewHander(ctx *fiber.Ctx) error {
	type article struct {
		Id string `json:"id"`
	}
	a := new(article)
	if err := ctx.BodyParser(a); err != nil {
		return err
	}
	article_ := models.ViewArticle(a.Id)

	tags := models.ViewArticlesTags(article_.Tags)
	article_.Tags = tags

	return ctx.JSON(fiber.Map{"data": article_, "status": true})
}
