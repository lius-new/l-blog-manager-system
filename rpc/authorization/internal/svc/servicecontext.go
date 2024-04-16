package svc

import (
	"github.com/lius-new/blog-backend/rpc/authorization/internal/config"
	model "github.com/lius-new/blog-backend/rpc/authorization/model/mongo"
	"github.com/lius-new/blog-backend/rpc/user/userer"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config config.Config
	Model  model.SecretModel
	Userer userer.Userer
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
		Model:  model.NewSecretModel(c.MongoURL, c.DBName, "secret", c.Cache),
		Userer: userer.NewUserer(zrpc.MustNewClient(c.User)),
	}
}
