package handler_admin

import (
	"github.com/lfxnxf/emo-frame/zd_http/server"
	"github.com/lfxnxf/emo-server/service"
)

var svc *service.Service

func Init(service *service.Service, g *server.HttpGroup) {
	svc = service
	initRouterAdmin(g)
}
