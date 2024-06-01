package main

import (
	"flag"
	"fmt"

	"github.com/lius-new/blog-backend/api/article/internal/config"
	"github.com/lius-new/blog-backend/api/article/internal/handler"
	"github.com/lius-new/blog-backend/api/article/internal/svc"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
)

var configFile = flag.String("f", "etc/article-api.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

  // TODO: 开发阶段设置跨域
  server := rest.MustNewServer(c.RestConf, rest.WithCors("http://localhost:5173"))
	defer server.Stop()

	ctx := svc.NewServiceContext(c)
	handler.RegisterHandlers(server, ctx)

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}
