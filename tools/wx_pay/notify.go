package wx_pay

import (
	"context"
	"github.com/lfxnxf/emo-frame/logging"
	"github.com/wechatpay-apiv3/wechatpay-go/core/auth/verifiers"
	"github.com/wechatpay-apiv3/wechatpay-go/core/downloader"
	"github.com/wechatpay-apiv3/wechatpay-go/core/notify"
	"go.uber.org/zap"
	"net/http"
)

func (w *WxPay) PayNotify(ctx context.Context, request *http.Request) (PayNotifyContent, error) {
	log := logging.For(ctx, "func", "PayNotify")

	certVisitor := downloader.MgrInstance().GetCertificateVisitor(w.config.MchID)

	handler := notify.NewNotifyHandler(w.config.MchAPIv3Key, verifiers.NewSHA256WithRSAVerifier(certVisitor))

	var content PayNotifyContent

	_, err := handler.ParseNotifyRequest(ctx, request, &content)
	if err != nil {
		log.Errorw("handler.ParseNotifyRequest error", zap.Error(err))
		return content, err
	}
	log.Infow("success!")
	return content, nil
}

// 退款回调
func (w *WxPay) RefundNotify(ctx context.Context, request *http.Request) (RefundNotifyContent, error) {
	log := logging.For(ctx, "func", "RefundNotify")

	certVisitor := downloader.MgrInstance().GetCertificateVisitor(w.config.MchID)

	handler := notify.NewNotifyHandler(w.config.MchAPIv3Key, verifiers.NewSHA256WithRSAVerifier(certVisitor))

	var content RefundNotifyContent

	_, err := handler.ParseNotifyRequest(ctx, request, &content)
	if err != nil {
		log.Errorw("handler.ParseNotifyRequest error", zap.Error(err))
		return content, err
	}
	log.Infow("success!")
	return content, nil
}
