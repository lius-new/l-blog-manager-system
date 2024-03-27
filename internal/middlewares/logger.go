package middlewares

import (
	"errors"
	"fmt"

	"github.com/gofiber/fiber/v2"
	Err "github.com/lius-new/liusnew-blog-backend-server/internal/errors"
	"github.com/lius-new/liusnew-blog-backend-server/internal/logger"
	"github.com/lius-new/liusnew-blog-backend-server/internal/models"
)

var NotRecordTraceIntercepts []string = []string{
	"/time",
}

func BaseLoggerMiddleware(ctx *fiber.Ctx) error {

	logInfoCotennt := fmt.Sprintf("[form]:%s:%s [target]:%s", ctx.IP(), ctx.Port(), ctx.Request().URI().FullURI())
	logger.Distribute(logger.LevelInfo, logInfoCotennt)

	path := string(ctx.Request().URI().Path())
	for _, v := range NotRecordTraceIntercepts {
		if v != path {
			models.Trace(ctx.IP(), path)
			break
		}
	}

	// 是否存在blocked, 如果在就组织
	if blockStatus := models.IsBlocked(ctx.IP()); blockStatus {
		return errors.New(Err.ErrorBlocked)
	}

	return ctx.Next()
}
