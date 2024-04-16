package main

import (
	"flag"
	"fmt"

	"github.com/lius-new/blog-backend/rpc/utils/internal/config"
	"github.com/lius-new/blog-backend/rpc/utils/internal/server"
	"github.com/lius-new/blog-backend/rpc/utils/internal/svc"
	"github.com/lius-new/blog-backend/rpc/utils/utils"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var configFile = flag.String("f", "etc/utils.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	ctx := svc.NewServiceContext(c)

	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		utils.RegisterUtilerServer(grpcServer, server.NewUtilerServer(ctx))

		if c.Mode == service.DevMode || c.Mode == service.TestMode {
			reflection.Register(grpcServer)
		}
	})
	defer s.Stop()

	fmt.Printf("Starting rpc server at %s...\n", c.ListenOn)
	s.Start()
}
