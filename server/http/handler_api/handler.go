package handler_api

import (
	"github.com/lfxnxf/emo-frame/zd_http/server"
	"github.com/lfxnxf/emo-server/service"
)

var svc *service.Service

const (
	ticket = "Oc6UwYSGhBUpaqzRlxfKVAPvsJECbvbT"
)

var notCheckSignUri = []string{
	"/api/v1/wechat/applet/upload_file",
}

func Init(service *service.Service, g *server.HttpGroup) {
	svc = service
	initRoute(g)
}
