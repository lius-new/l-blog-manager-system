package main

import (
	"github.com/lius-new/liusnew-blog-backend-server/internal/routers"
)

func main() {
	go routers.Server1()
	go routers.Server2()
	select {}
}
