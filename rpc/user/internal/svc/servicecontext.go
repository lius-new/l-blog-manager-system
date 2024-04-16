package svc

import (
	"github.com/lius-new/blog-backend/rpc/user/internal/config"
	model "github.com/lius-new/blog-backend/rpc/user/model/mongo"
	"github.com/lius-new/blog-backend/rpc/utils/utiler"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config config.Config
	Model  model.UserModel
	Utiler utiler.Utiler
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
		Utiler: utiler.NewUtiler(zrpc.MustNewClient(c.Utils)),
		Model:  model.NewUserModel(c.MongoURL, c.DBName, "user", c.Cache),
	}
}
