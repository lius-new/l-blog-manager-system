package middlewares

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/lius-new/liusnew-blog-backend-server/internal/logger"
)

func BaseLoggerMiddleware(c *fiber.Ctx) error {
	logInfoCotennt := fmt.Sprintf("[form]:%s:%s [target]:%s", c.IP(), c.Port(), c.Request().URI().FullURI())

	logger.Distribute(logger.LevelInfo, logInfoCotennt)

	return c.Next()
}
