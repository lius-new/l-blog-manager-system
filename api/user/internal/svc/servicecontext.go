package svc

import (
	"github.com/lius-new/blog-backend/api/user/internal/config"
	"github.com/lius-new/blog-backend/api/user/internal/middleware"
	"github.com/lius-new/blog-backend/rpc/analyzer/analyzerclient"
	"github.com/lius-new/blog-backend/rpc/authorization/auther"
	"github.com/lius-new/blog-backend/rpc/user/userer"
	"github.com/lius-new/blog-backend/rpc/utils/utiler"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config             config.Config
	Auther             auther.Auther
	Userer             userer.Userer
	Utiler             utiler.Utiler
	AuthMiddleware     rest.Middleware
	AnalyzerMiddleware rest.Middleware
}

func NewServiceContext(c config.Config) *ServiceContext {
	auther := auther.NewAuther(zrpc.MustNewClient(c.Authorization))
	analyzer := analyzerclient.NewAnalyzer(zrpc.MustNewClient(c.Analyzer))

	return &ServiceContext{
		Config:             c,
		Auther:             auther,
		Userer:             userer.NewUserer(zrpc.MustNewClient(c.User)),
		Utiler:             utiler.NewUtiler(zrpc.MustNewClient(c.Utils)),
		AuthMiddleware:     middleware.NewAuthMiddleware(auther).Handle,
		AnalyzerMiddleware: middleware.NewAnalyzerMiddleware(analyzer).Handle,
	}
}
