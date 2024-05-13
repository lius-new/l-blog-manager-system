package tests

import (
	"github.com/zeromicro/go-zero/core/conf"

	"github.com/lius-new/blog-backend/rpc/analyzer/internal/config"
	"github.com/lius-new/blog-backend/rpc/analyzer/internal/svc"
)

var SVC_CONTEXT *svc.ServiceContext

func init() {
	configFile := "../../../etc/content.yaml"

	var c config.Config
	conf.MustLoad(configFile, &c)
	SVC_CONTEXT = svc.NewServiceContext(c)
}
