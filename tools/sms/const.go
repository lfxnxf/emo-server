package sms

const (
	endpoint = "dysmsapi.aliyuncs.com"

	// SignMetaSports 签名
	SignMetaSports = "MetaSports"

	TemplatePublicValidateCode                  = "SMS_243035087" // 通用验证码
	TemplatePriceOrderRefund                    = "SMS_244175032" // 普通运营订单（定价购买）退款
	TemplateCompetitionOrderRefund              = "SMS_249225840" // 赛事取消退款
	TemplateCourseCancel                        = "SMS_255300594" // 课程取消退还现金
	TemplateCourseCancelSendBackVitalityCard    = "SMS_255385595" // 课程取消退还元气卡
	TemplateCourseRefund                        = "SMS_255280612" // 课程售后退款
	TemplateCourseQueueRefund                   = "SMS_255245622" // 课程排队退款
	TemplateVitalityCardOrderRefund             = "SMS_253800036" // 元气卡订单退款
	TemplateOrderCancelSendBackVitalityCard     = "SMS_255195560" // 课程售后退元气卡
	TemplateQueueCancelSendBackVitalityCard     = "SMS_255400616" // 课程排队退元气卡
	TemplateBuyPrivateCourse                    = "SMS_267655556" // 购买私教课程
	TemplatePrivateCourseRefund                 = "SMS_255280612" // 私教课程退款，使用课程售后退款相同模板
	TemplatePrivateCourseReserveToUser          = "SMS_267900528" // 私教课程预约-客户端
	TemplatePrivateCourseReserveToTeacher       = "SMS_267675569" // 私教课程预约-教练端
	TemplatePrivateCourseReserveCancelToUser    = "SMS_267860511" // 私教课程取消预约-客户端
	TemplatePrivateCourseReserveCancelToTeacher = "SMS_267855470" // 私教课程取消预约-教练
)
