package wx_pay

import (
	"context"
	"github.com/lfxnxf/emo-frame/logging"
	"github.com/lfxnxf/emo-frame/utils"
	"github.com/wechatpay-apiv3/wechatpay-go/core"
	"github.com/wechatpay-apiv3/wechatpay-go/services/payments/native"
	"go.uber.org/zap"
	"time"
)

func (w *WxPay) CreateNativeOrder(ctx context.Context, data NativeCreateOrder) (string, error) {
	log := logging.For(ctx, "func", "CreateNativeOrder",
		zap.Any("data", data),
	)

	svc := native.NativeApiService{Client: w.client}
	resp, _, err := svc.Prepay(ctx,
		native.PrepayRequest{
			Appid:       core.String(w.config.AppId),
			Mchid:       core.String(w.config.MchID),
			Description: core.String(data.Description),
			OutTradeNo:  core.String(data.OutTradeNo),
			TimeExpire:  core.Time(time.Now()),
			Attach:      core.String(data.Attach),
			NotifyUrl:   core.String(w.config.PayNotifyUrl),
			Amount: &native.Amount{
				Currency: core.String("CNY"),
				Total:    core.Int64(data.Amount),
			},
		},
	)

	if err != nil {
		// 处理错误
		log.Errorw("native.NativeApiService.Prepay error", zap.Error(err))
		return "", err
	}

	return utils.StringPointerVal(resp.CodeUrl), nil
}
