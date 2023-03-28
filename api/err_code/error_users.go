package err_code

import "github.com/lfxnxf/emo-frame/zd_error"

// users 1001开头

var (
	ErrorUserNotLogin = zd_error.AddError(1001000, "用户未登录")
)
