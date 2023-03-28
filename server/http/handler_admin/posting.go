package handler_admin

import (
	"github.com/lfxnxf/emo-frame/logging"
	"github.com/lfxnxf/emo-frame/zd_http/http_ctx"
	"github.com/lfxnxf/emo-server/model/model_posting"
	"go.uber.org/zap"
)

// 导出设备表格
func addPosting(c *http_ctx.HttpContext) {
	log := logging.For(c, "func", "addPosting")

	var req model_posting.AddPostingReq
	err := c.ShouldBindJSON(&req)
	if err != nil {
		log.Errorw("params error", zap.Error(err))
		c.WriteParamsError(err, &req)
		return
	}
	return
}

