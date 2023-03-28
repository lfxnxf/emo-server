package model

type WechatPayReply struct {
	PrepayId  string `json:"prepay_id"`
	AppId     string `json:"app_id"`
	TimeStamp string `json:"time_stamp"`
	NonceStr  string `json:"nonce_str"`
	Package   string `json:"package"`
	SignType  string `json:"sign_type"`
	PaySign   string `json:"pay_sign"`
}
