package svc

import (
	"github.com/lius-new/blog-backend/api/article/internal/config"
	"github.com/lius-new/blog-backend/api/article/internal/middleware"
	"github.com/lius-new/blog-backend/rpc/authorization/auther"
	"github.com/lius-new/blog-backend/rpc/content/contentclient"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config         config.Config
	Auth           auther.Auther
	Content        contentclient.Content
	AuthMiddleware rest.Middleware
}

func NewServiceContext(c config.Config) *ServiceContext {
	auther := auther.NewAuther(zrpc.MustNewClient(c.Authorization))

	return &ServiceContext{
		Config:         c,
		Auth:           auther,
		Content:        contentclient.NewContent(zrpc.MustNewClient(c.Content)),
		AuthMiddleware: middleware.NewAuthMiddleware(auther).Handle,
	}
}
