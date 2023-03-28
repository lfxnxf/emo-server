package http

import (
	"github.com/lfxnxf/emo-frame/inits"
	"github.com/lfxnxf/emo-frame/logging"
	"github.com/lfxnxf/emo-frame/zd_http/http_ctx"
	"github.com/lfxnxf/emo-frame/zd_http/server"
	"github.com/lfxnxf/emo-server/server/http/handler_admin"
	"github.com/lfxnxf/emo-server/server/http/handler_api"
	"github.com/lfxnxf/emo-server/service"
)

var (
	httpServer *server.HttpServer
	svc        *service.Service
)

// Init create a http server and run it
func Init(s *service.Service) {
	svc = s

	// new http server
	httpServer = inits.NewHttpServer()

	httpServer.Any("/ping", func(c *http_ctx.HttpContext) {
		c.JSON(200, "pong")
	})

	// 客户端接口
	apiGroup := httpServer.Group("/api")
	handler_api.Init(s, apiGroup)

	// 后台接口
	adminGroup := httpServer.Group("/api/admin")
	handler_admin.Init(s, adminGroup)

	// start a http server
	go func() {
		if err := httpServer.Start(); err != nil {
			logging.Fatalf("http server start failed, err %v", err)
		}
	}()
}

func Shutdown() {
	httpServer.Shutdown()
}
