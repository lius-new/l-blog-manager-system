package middlewares

import (
	"errors"
	"os"
	"strings"

	"github.com/gofiber/fiber/v2"
	Err "github.com/lius-new/liusnew-blog-backend-server/internal/errors"
	"github.com/lius-new/liusnew-blog-backend-server/internal/jwt"
	"github.com/lius-new/liusnew-blog-backend-server/internal/logger"
)

var NotIntercepts []string = []string{
	"/api/user/login", "/api/articles/view", "/api/articles/views", "/api/articles/search", "/time", "/favicon.ico", "/api/file/:hash", "/",
}

func AuthMiddleware(c *fiber.Ctx) error {
	path := string(c.Request().URI().Path())

	for _, v := range NotIntercepts {
		if v == path {
			return c.Next()
		} else if strings.Contains(v, ":") {
			splits := strings.Split(v, ":")
			if strings.HasPrefix(path, splits[0]) {
				return c.Next()
			}
		}
	}

	secret := c.Request().Header.Cookie("secret")

	tokenClaims, err := jwt.JWT.ParseJwtToken(os.Getenv("SECRET_VALUE_2"), string(secret))
	if err != nil {
		c.ClearCookie()
		logger.Panic("AuthPanic", err.Error())
		return errors.New(Err.ErrorUnauthorized)
	}

	if _, err := jwt.JWT.ParseJwtToken(os.Getenv("SECRET_VALUE"), tokenClaims.Token); err != nil {
		c.ClearCookie()
		logger.Panic("AuthPanic", err.Error())
		return errors.New(Err.ErrorUnauthorized)
	}

	return c.Next()
}
