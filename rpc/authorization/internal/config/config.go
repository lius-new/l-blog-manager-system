package config

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	zrpc.RpcServerConf
	MongoURL string
	DBName   string
	Cache    cache.CacheConf
	User     zrpc.RpcClientConf
}
