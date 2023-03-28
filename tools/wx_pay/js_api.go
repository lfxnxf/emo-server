package wx_pay

import (
	"context"
	"github.com/lfxnxf/emo-frame/logging"
	"github.com/lfxnxf/emo-frame/utils"
	"github.com/lfxnxf/emo-server/model"
	"github.com/wechatpay-apiv3/wechatpay-go/core"
	"github.com/wechatpay-apiv3/wechatpay-go/services/payments/jsapi"
	"go.uber.org/zap"
)

// js_api统一下单
func (w *WxPay) JsApiPayCreateOrder(ctx context.Context, data JsApiCreateOrder) (model.WechatPayReply, error) {
	log := logging.For(ctx, "func", "JsApiPayCreateOrder",
		zap.Any("data", data),
	)
	svc := jsapi.JsapiApiService{Client: w.client}

	// 得到prepay_id，以及调起支付所需的参数和签名
	resp, _, err := svc.PrepayWithRequestPayment(ctx,
		jsapi.PrepayRequest{
			Appid:       core.String(w.config.AppId),
			Mchid:       core.String(w.config.MchID),
			Description: core.String(data.Description),
			OutTradeNo:  core.String(data.OutTradeNo),
			Attach:      core.String(data.Attach),
			NotifyUrl:   core.String(w.config.PayNotifyUrl),
			Amount: &jsapi.Amount{
				Total: core.Int64(data.Amount),
			},
			Payer: &jsapi.Payer{
				Openid: core.String(data.Openid),
			},
		},
	)

	if err != nil {
		log.Errorw("svc.PrepayWithRequestPayment error", zap.Error(err))
		return model.WechatPayReply{}, err
	}

	log.Infow("success!", zap.Any("resp", resp))

	return model.WechatPayReply{
		PrepayId:  utils.StringPointerVal(resp.PrepayId),
		AppId:     utils.StringPointerVal(resp.Appid),
		TimeStamp: utils.StringPointerVal(resp.TimeStamp),
		NonceStr:  utils.StringPointerVal(resp.NonceStr),
		Package:   utils.StringPointerVal(resp.Package),
		SignType:  utils.StringPointerVal(resp.SignType),
		PaySign:   utils.StringPointerVal(resp.PaySign),
	}, nil
}

// 调用支付单信息
func (w *WxPay) GetJsApiPayResult(ctx context.Context, outTradeNo string) (JsApiPayResult, error) {
	log := logging.For(ctx, "func", "GetJsApiPayResult",
		zap.String("out_trade_no", outTradeNo),
	)

	svc := jsapi.JsapiApiService{Client: w.client}
	resp, _, err := svc.QueryOrderByOutTradeNo(ctx,
		jsapi.QueryOrderByOutTradeNoRequest{
			OutTradeNo: core.String(outTradeNo),
			Mchid:      core.String(w.config.MchID),
		},
	)
	if err != nil {
		log.Errorw("svc.QueryOrderByOutTradeNo error", zap.Error(err))
		return JsApiPayResult{}, err
	}

	return JsApiPayResult{
		Amount:         utils.Int64PointerVal(resp.Amount.Total),
		AppId:          utils.StringPointerVal(resp.Appid),
		Attach:         utils.StringPointerVal(resp.Attach),
		BankType:       utils.StringPointerVal(resp.BankType),
		MchId:          utils.StringPointerVal(resp.Mchid),
		OutTradeNo:     utils.StringPointerVal(resp.OutTradeNo),
		SuccessTime:    utils.StringPointerVal(resp.SuccessTime),
		TradeState:     utils.StringPointerVal(resp.TradeState),
		TradeStateDesc: utils.StringPointerVal(resp.TradeStateDesc),
		TradeType:      utils.StringPointerVal(resp.TradeType),
		TransactionId:  utils.StringPointerVal(resp.TransactionId),
	}, nil
}
