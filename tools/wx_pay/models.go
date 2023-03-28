package wx_pay

import "time"

const (
	TradeStateSuccess   = "SUCCESS"    // 支付成功
	TradeStateRefund    = "REFUND"     // 转入退款
	TradeStateNoPay     = "NOTPAY"     // 未支付
	TradeStateClosed    = "CLOSED"     // 已关闭
	TradeStateRevoked   = "REVOKED"    // 已撤销(刷卡支付)
	TradeStatePaying    = "USERPAYING" // 用户支付中
	TradeStatePayError  = "PAYERROR"   // 支付失败(其他原因，如银行返回失败)
	TradeStatePayAccept = "ACCEPT"     // 已接收，等待扣款
)

type JsApiPayResult struct {
	Amount         int64  `json:"amount"`
	AppId          string `json:"app_id"`
	Attach         string `json:"attach"`
	BankType       string `json:"bank_type"`
	MchId          string `json:"mch_id"`
	OutTradeNo     string `json:"out_trade_no"`
	SuccessTime    string `json:"success_time"`
	TradeState     string `json:"trade_state"`
	TradeStateDesc string `json:"trade_state_desc"`
	TradeType      string `json:"trade_type"`
	TransactionId  string `json:"transaction_id"`
}

type JsApiCreateOrder struct {
	Description string
	OutTradeNo  string
	Attach      string
	Amount      int64
	Openid      string
}

type PayAttach struct {
	OrderType int64 `json:"order_type"`
}

type NativeCreateOrder struct {
	Description string
	OutTradeNo  string
	Attach      string
	Amount      int64
}

// 退款单申请
type CreateRefundOrder struct {
	OutTradeNo  string `json:"out_trade_no"`
	OutRefundNo string `json:"out_refund_no"`
	Reason      string `json:"reason"`
	Amount      int64  `json:"amount"`
	Total       int64  `json:"total"`
}

const (
	RefundStatusSuccess    = "SUCCESS"    // 退款成功
	RefundStatusClosed     = "CLOSED"     // 退款关闭
	RefundStatusProcessing = "PROCESSING" // 退款中
	RefundStatusAbnormal   = "ABNORMAL"   // 退款异常，退款到银行发现用户的卡作废或者冻结了，导致原路退款银行卡失败，可前往【商户平台—>交易中心】，手动处理此笔退款
)

type RefundOrderReply struct {
	RefundId      string    `json:"refund_id"`
	OutRefundNo   string    `json:"out_refund_no"`
	TransactionId string    `json:"transaction_id"`
	OutTradeNo    string    `json:"out_trade_no"`
	SuccessTime   time.Time `json:"success_time"`
	CreateTime    time.Time `json:"create_time"`
	Status        string    `json:"status"`
	Total         int64     `json:"total"`
	RefundAmount  int64     `json:"refund_amount"`
}

// 退款回调
type RefundNotifyContent struct {
	MchId         string    `json:"mchid"`
	OutTradeNo    string    `json:"out_trade_no"`
	TransactionID string    `json:"transaction_id"`
	OutRefundNo   string    `json:"out_refund_no"`
	RefundID      string    `json:"refund_id"`
	RefundStatus  string    `json:"refund_status"`
	SuccessTime   time.Time `json:"success_time"`
	Amount        struct {
		Total       int64 `json:"total"`
		Refund      int64 `json:"refund"`
		PayerTotal  int64 `json:"payer_total"`
		PayerRefund int64 `json:"payer_refund"`
	} `json:"amount"`
	UserReceivedAccount string `json:"user_received_account"`
}

// 支付回调
type PayNotifyContent struct {
	MchId          string    `json:"mchid"`
	AppId          string    `json:"appid"`
	OutTradeNo     string    `json:"out_trade_no"`
	TransactionID  string    `json:"transaction_id"`
	TradeType      string    `json:"trade_type"`
	TradeState     string    `json:"trade_state"`
	TradeStateDesc string    `json:"trade_state_desc"`
	BankType       string    `json:"bank_type"`
	Attach         string    `json:"attach"`
	SuccessTime    time.Time `json:"success_time"`
	Payer          struct {
		Openid string `json:"openid"`
	} `json:"payer"`
	Amount struct {
		Total         int64  `json:"total"`
		PayerTotal    int64  `json:"payer_total"`
		Currency      string `json:"currency"`
		PayerCurrency string `json:"payer_currency"`
	} `json:"amount"`
}
