package tests

import (
	"github.com/lius-new/blog-backend/api/article/internal/config"
	"github.com/lius-new/blog-backend/api/article/internal/svc"
	"github.com/zeromicro/go-zero/core/conf"
)

var SVC_CONTEXT *svc.ServiceContext

func init() {
	var configFile = "../../../etc/article-api.yaml"
	var c config.Config
	conf.MustLoad(configFile, &c)
	SVC_CONTEXT = svc.NewServiceContext(c)
}
