package tests

import (
	"github.com/lius-new/blog-backend/rpc/utils/internal/config"
	"github.com/lius-new/blog-backend/rpc/utils/internal/svc"
)

var SVC_CONTEXT *svc.ServiceContext

func init() {
	var c config.Config
	// c.MongoURL = "mongodb://lius:lsmima@127.0.0.1:27017"
	// c.DBName = "liusnew-blog"
	// c.Cache = cache.ClusterConf{cache.NodeConf{Weight: 10, RedisConf: redis.RedisConf{Host: "127.0.0.1:6379", Type: "node"}}}
	SVC_CONTEXT = svc.NewServiceContext(c)
}
