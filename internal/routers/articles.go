package routers

import (
	"fmt"
	"mime/multipart"
	"os"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/lius-new/liusnew-blog-backend-server/internal/models"
	"github.com/lius-new/liusnew-blog-backend-server/internal/utils"
)

func RegisterArticlesHanlder(app *fiber.App) {
	api := app.Group("/api/articles")

	api.Post("/create", createHander)
	api.Post("/modify", modifyHander)
	api.Post("/modify-status", modifyHanderStatus)
	api.Delete("/delete", deleteHander)
	api.Post("/views", viewsHander)
	api.Post("/view", viewHander)
}

func RegisterArticlesHanlder2(app *fiber.App) {
	api := app.Group("/api/articles")

	api.Post("/views", viewsHander2)
	api.Post("/view", viewHander)
	api.Post("/search", searchHander2)
}

func createHander(ctx *fiber.Ctx) error {
	var (
		form *multipart.Form
		err  error
	)

	if form, err = ctx.MultipartForm(); err != nil {
		return err
	}

	var title, content string
	var tags []string

	if tempValue := form.Value["title"]; len(tempValue) > 0 {
		title = tempValue[0]
	}
	if tempValue := form.Value["content"]; len(tempValue) > 0 {
		content = tempValue[0]
	}
	if tempValue := form.Value["tags"]; len(tempValue) > 0 {
		tags = tempValue
	}

	// 文件
	files := form.File["covers"]

	if len(title) == 0 || len(content) == 0 || len(tags) == 0 {
		return ctx.SendStatus(fiber.ErrBadRequest.Code)
	}

	covers := make([]string, 0)
	for _, v := range files {
		file, err := v.Open()
		if err != nil {
			return err
		}
		defer file.Close()
		hash, err := utils.Hash(file)
		if err != nil {
			return err
		}
		covers = append(covers, hash)
	}

	tags = models.SaveTags(tags)
	article_, err := models.CreateArticles(title, content, tags, covers)

	if err != nil && err.Error() == "article exist" {
		return ctx.JSON(fiber.Map{"message": "article exist"})
	}

	article_.Tags = tags

	// 保存文件到本地, 为什么不放在上面for一起呢? 因为有可能保存失败
	for i, file := range files {
		// save to example : COVER_PATH/文件hash.后缀 =>./data/covers/hsdjf24hjsfh283sf.png
		savePath := fmt.Sprintf("%s/%s.%s", os.Getenv("COVER_PATH"), covers[i], utils.GetFileSuffix(file.Filename))
		if exist, _ := utils.FileExist(savePath, os.Getenv("COVER_PATH")); len(exist) == 0 { // 不存在就保存
			if err := ctx.SaveFile(file, savePath); err != nil {
				return err
			}
		}
	}

	return ctx.JSON(fiber.Map{"data": article_, "status": true})
}

func modifyHanderStatus(ctx *fiber.Ctx) error {
	param := struct {
		Id     string `json:"id"`
		Status bool   `json:"status"`
	}{}

	if err := ctx.BodyParser(&param); err != nil {
		return err
	}
	article, err := models.ModifyArticleStatus(param.Id, param.Status)

	if err != nil && err.Error() == "article not found" {
		return ctx.SendStatus(fiber.ErrNotFound.Code)
	}

	return ctx.JSON(fiber.Map{"data": article, "status": true})
}
func modifyHander(ctx *fiber.Ctx) error {
	var (
		form *multipart.Form
		err  error
	)

	if form, err = ctx.MultipartForm(); err != nil {
		return err
	}

	var id, title, content string
	var tags []string
	var status bool

	if tempValue := form.Value["id"]; len(tempValue) > 0 {
		id = tempValue[0]
	}
	if tempValue := form.Value["title"]; len(tempValue) > 0 {
		title = tempValue[0]
	}
	if tempValue := form.Value["content"]; len(tempValue) > 0 {
		content = tempValue[0]
	}
	if tempValue := form.Value["tags"]; len(tempValue) > 0 {
		tags = tempValue
	}
	if tempValue := form.Value["status"]; len(tempValue) > 0 {
		if tempValue[0] == "false" {
			status = false
		} else {
			status = true
		}
	}

	// 文件
	files := form.File["covers"]

	if len(id) == 0 {
		return ctx.SendStatus(fiber.ErrBadRequest.Code)
	}

	covers := make([]string, 0)
	for _, v := range files {
		file, err := v.Open()
		if err != nil {
			return err
		}
		defer file.Close()
		hash, err := utils.Hash(file)
		if err != nil {
			return err
		}
		covers = append(covers, hash)
	}
	if tempValue := form.Value["covers"]; len(tempValue) > 0 {
		covers = append(covers, tempValue...)
	}
	tags = models.SaveTags(tags)

	article_, err := models.ModifyArticles(id, title, content, tags, covers, status)
	if err != nil && err.Error() == "article not found" {
		return ctx.SendStatus(fiber.ErrNotFound.Code)
	}

	tags = models.ViewArticlesTags(article_.Tags)
	article_.Tags = tags

	// 保存文件到本地, 为什么不放在上面for一起呢? 因为有可能保存失败
	for i, file := range files {
		// save to example : COVER_PATH/文件hash.后缀 =>./data/covers/hsdjf24hjsfh283sf.png
		savePath := fmt.Sprintf("%s/%s.%s", os.Getenv("COVER_PATH"), covers[i], utils.GetFileSuffix(file.Filename))
		if exist, _ := utils.FileExist(savePath, os.Getenv("COVER_PATH")); len(exist) == 0 { // 不存在就保存
			if err := ctx.SaveFile(file, savePath); err != nil {
				return err
			}
		}
	}

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

	articles, count := models.ViewArticles(a.PageSize, a.PageNum, true)

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
func viewsHander2(ctx *fiber.Ctx) error {
	type article struct {
		PageSize int64 `json:"page_size"`
		PageNum  int64 `json:"page_num"`
	}

	a := new(article)
	if err := ctx.BodyParser(a); err != nil {
		return err
	}

	articles, count := models.ViewArticles(a.PageSize, a.PageNum, false)

	// 获取每个文章的标签名
	for index := range articles {
		tags := models.ViewArticlesTags(articles[index].Tags)
		articles[index].Tags = tags
		articles[index].Covers = make([]string, 0)
		if len(articles[index].Content) > 300 {
			articles[index].Content = strings.Join([]string{articles[index].Content[:300], "..."}, "")
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

func searchHander2(ctx *fiber.Ctx) error {
	type article struct {
		Title string `json:"title"`
	}
	a := new(article)
	if err := ctx.BodyParser(a); err != nil {
		return err
	}
	articles := models.SearchArticle(a.Title)
	for index := range articles {
		articles[index].Tags = make([]string, 0)
		articles[index].Covers = make([]string, 0)
		articles[index].Content = ""
	}

	return ctx.JSON(fiber.Map{"data": articles, "status": true})
}
