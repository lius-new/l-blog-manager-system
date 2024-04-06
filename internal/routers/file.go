package routers

import (
	"fmt"
	"mime/multipart"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/lius-new/liusnew-blog-backend-server/internal/utils"
)

func RegisterFileHanlder(app *fiber.App) {
	api := app.Group("/api/file")
	api.Get("/:hash", viewFileHandle)
	api.Post("/upload-images", uploadImageFileHandle)
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

// uploadImageFileHandle: 保存图片
func uploadImageFileHandle(ctx *fiber.Ctx) error {
	var (
		form *multipart.Form
		err  error
	)

	if form, err = ctx.MultipartForm(); err != nil {
		return err
	}
	files := form.File["images"]

	hashs := make([]string, 0)
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
		hashs = append(hashs, hash)
	}

	for i, file := range files {
		// save to example : COVER_PATH/文件hash.后缀 =>./data/covers/hsdjf24hjsfh283sf.png
		savePath := fmt.Sprintf("%s/%s.%s", os.Getenv("COVER_PATH"), hashs[i], utils.GetFileSuffix(file.Filename))
		if exist, _ := utils.FileExist(savePath, os.Getenv("COVER_PATH")); len(exist) == 0 { // 不存在就保存
			if err := ctx.SaveFile(file, savePath); err != nil {
				return err
			}
		}
	}

	return ctx.JSON(fiber.Map{"status": true, "data": hashs})
}
