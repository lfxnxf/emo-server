package err_code

import "github.com/lfxnxf/emo-frame/zd_error"

var (
	ErrorOperateOften = zd_error.AddError(1000000, "操作频繁，请稍后再试")
)

