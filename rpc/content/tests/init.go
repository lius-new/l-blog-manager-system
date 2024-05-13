package tests

import (
	"github.com/lius-new/blog-backend/rpc/content/internal/config"
	"github.com/lius-new/blog-backend/rpc/content/internal/svc"
	"github.com/zeromicro/go-zero/core/conf"
)

var SVC_CONTEXT *svc.ServiceContext

func init() {

	var configFile = "../../../etc/content.yaml"

	var c config.Config
	conf.MustLoad(configFile, &c)
	SVC_CONTEXT = svc.NewServiceContext(c)
}
