package tests

import (
	"github.com/lius-new/blog-backend/rpc/user/internal/config"
	"github.com/lius-new/blog-backend/rpc/user/internal/svc"
	"github.com/zeromicro/go-zero/core/conf"
)

var SVC_CONTEXT *svc.ServiceContext

func init() {

	var configFile = "../../../etc/user.yaml"

	var c config.Config
	conf.MustLoad(configFile, &c)
	SVC_CONTEXT = svc.NewServiceContext(c)
}
