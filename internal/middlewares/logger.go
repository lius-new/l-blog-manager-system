package middlewares

import (
	"errors"
	"fmt"

	"github.com/gofiber/fiber/v2"
	Err "github.com/lius-new/liusnew-blog-backend-server/internal/errors"
	"github.com/lius-new/liusnew-blog-backend-server/internal/logger"
	"github.com/lius-new/liusnew-blog-backend-server/internal/models"
)

// 该链接不记录. 资源白名单
var NotRecordPathTraceIntercepts []string = []string{
	"/time",
}

// 该IP不记录不记录. Ip白名单
var NotRecordIPTraceIntercepts []string = []string{
	"101.42.173.48",
}

func BaseLoggerMiddleware(ctx *fiber.Ctx) error {
	logInfoCotennt := fmt.Sprintf("[form]:%s:%s [target]:%s", ctx.IP(), ctx.Port(), ctx.Request().URI().FullURI())
	logger.Distribute(logger.LevelInfo, logInfoCotennt)

	// 是否存在blocked, 如果在就阻止
	if blockStatus := models.IsBlocked(ctx.IP()); blockStatus {
		return errors.New(Err.ErrorBlocked)
	}

	path := string(ctx.Request().URI().Path())
	if !isExist(path, NotRecordPathTraceIntercepts) || !isExist(ctx.IP(), NotRecordIPTraceIntercepts) {
		models.Trace(ctx.IP(), path)
	}

	return ctx.Next()
}

// 判断是否存在
func isExist(current string, array []string) bool {
	for _, v := range array {
		if v == current {
			return true
		}
	}
	return false
}
