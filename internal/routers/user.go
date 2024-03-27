package routers

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/lius-new/liusnew-blog-backend-server/internal/jwt"
	"github.com/lius-new/liusnew-blog-backend-server/internal/models"
)

func RegisterUserHanlder(app *fiber.App) {
	api := app.Group("/api/user")

	api.Post("/login", loginHandler)

}

func loginHandler(c *fiber.Ctx) error {
	type user struct {
		Username string `json:"username" bind:"required"`
		Password string `json:"password" bind:"required"`
	}
	u := new(user)
	if err := c.BodyParser(u); err != nil {
		log.Println(err)
		return err
	}

	res, err := models.Login(u.Username, u.Password)

	if err != nil || len(res) != 2 {
		return c.JSON(fiber.Map{
			"message": "Login Failed",
		})
	}

	token, err := jwt.JWT.GenerateJwtToken(res[0], res[1], os.Getenv("SECRET_VALUE"), "", jwt.JWT.GetExpiresAt())
	if err != nil {
		panic(err)
	}
	tokenSecond, err := jwt.JWT.GenerateJwtTokenSecond(os.Getenv("SECRET_VALUE_2"), "", jwt.JWT.GetExpiresAt(), res[1], token)
	if err != nil {
		panic(err)
	}

	c.Cookie(&fiber.Cookie{
		Name:    "secret",
		Value:   tokenSecond,
		Expires: jwt.JWT.GetExpiresAt(),
	})

	return c.JSON(fiber.Map{
		"message": "Login Successed",
	})
}
