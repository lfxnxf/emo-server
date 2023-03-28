package handler_admin

import (
	"github.com/lfxnxf/emo-frame/zd_http/server"
)

func initRouterAdmin(g *server.HttpGroup) {

	// 帖子
	postingGroup := g.Group("/model_posting")
	{
		postingGroup.POST("/add", addPosting)
	}
}
