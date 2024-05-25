package config

import (
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	rest.RestConf
	Authorization zrpc.RpcClientConf
	User          zrpc.RpcClientConf
	Utils         zrpc.RpcClientConf
	Analyzer      zrpc.RpcClientConf
}
