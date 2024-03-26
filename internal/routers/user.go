package routers

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/lius-new/liusnew-blog-backend-server/internal/jwt"
	"github.com/lius-new/liusnew-blog-backend-server/internal/models"
)

func RegisterUserHanlder(app *fiber.App) {
	api := app.Group("/api/user")

	api.Post("/user/login", func(c *fiber.Ctx) error {
		type user struct {
			Username string `json:"username" bind:"required"`
			Password string `json:"password" bind:"required"`
		}
		u := new(user)
		if err := c.BodyParser(u); err != nil {
			log.Println(err)
			return err
		}

		res := models.Login(u.Username, u.Password)
		token, err := jwt.JWT.GenerateJwtToken(res[0], res[1], "", "", jwt.JWT.GetExpiresAt())
		if err != nil {
			panic(err)
		}

		return c.JSON(fiber.Map{
			"token": token,
		})
	})

}
