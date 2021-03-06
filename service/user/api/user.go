package main

import (
	"flag"
	"fmt"
	"mall/common/middleware"
	"mall/service/user/api/internal/config"
	"mall/service/user/api/internal/handler"
	"mall/service/user/api/internal/svc"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
)

var configFile = flag.String("f", "etc/user.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	ctx := svc.NewServiceContext(c)
	server := rest.MustNewServer(c.RestConf)
	defer server.Stop()

	// 全局中间件
	//将网关验证后的userId设置到ctx中
	server.Use(middleware.NewSetUidToCtxMiddleware().Handle)
	server.Use(middleware.LoggerMiddleware)
	handler.RegisterHandlers(server, ctx)

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}
