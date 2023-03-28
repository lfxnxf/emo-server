package wx_pay

import (
	"context"
	"github.com/lfxnxf/emo-frame/logging"
	"github.com/lfxnxf/emo-frame/utils"
	"github.com/wechatpay-apiv3/wechatpay-go/core"
	"github.com/wechatpay-apiv3/wechatpay-go/services/refunddomestic"
	"go.uber.org/zap"
)

// 申请微信退款单
func (w *WxPay) CreateRefundOrder(ctx context.Context, data CreateRefundOrder) (RefundOrderReply, error) {
	log := logging.For(ctx, "func", "CreateRefundOrder",
		zap.Any("data", data),
	)

	svc := refunddomestic.RefundsApiService{Client: w.client}

	// 申请退款单
	resp, _, err := svc.Create(ctx,
		refunddomestic.CreateRequest{
			OutRefundNo:  core.String(data.OutRefundNo),
			OutTradeNo:   core.String(data.OutTradeNo),
			Reason:       core.String(data.Reason),
			NotifyUrl:    core.String(w.config.RefundNotifyUrl),
			FundsAccount: refunddomestic.REQFUNDSACCOUNT_AVAILABLE.Ptr(),
			Amount: &refunddomestic.AmountReq{
				Currency: core.String("CNY"),
				Refund:   core.Int64(data.Amount),
				Total:    core.Int64(data.Total),
			},
		},
	)

	if err != nil {
		// 处理错误
		log.Errorw("refund create err", zap.Error(err))
		return RefundOrderReply{}, err
	}

	return w.getRefundResult(resp), nil
}

// 获取退款结果
func (w *WxPay) GetRefundResult(ctx context.Context, outRefundNo string) (RefundOrderReply, error) {
	log := logging.For(ctx, "func", "GetRefundResult",
		zap.String("out_refund_no", outRefundNo),
	)
	svc := refunddomestic.RefundsApiService{Client: w.client}
	resp, _, err := svc.QueryByOutRefundNo(ctx,
		refunddomestic.QueryByOutRefundNoRequest{
			OutRefundNo: core.String("1217752501201407033233368018"),
			SubMchid:    core.String("1900000109"),
		},
	)

	if err != nil {
		// 处理错误
		log.Errorw("get refund result", zap.Error(err))
		return RefundOrderReply{}, err
	}

	return w.getRefundResult(resp), nil
}

func (w *WxPay) getRefundResult(resp *refunddomestic.Refund) RefundOrderReply {
	var status string
	if resp.Status.Ptr() != nil {
		status = string(*resp.Status.Ptr())
	}
	return RefundOrderReply{
		RefundId:      utils.StringPointerVal(resp.RefundId),
		OutRefundNo:   utils.StringPointerVal(resp.OutRefundNo),
		TransactionId: utils.StringPointerVal(resp.TransactionId),
		OutTradeNo:    utils.StringPointerVal(resp.OutTradeNo),
		SuccessTime:   utils.TimePointerVal(resp.SuccessTime),
		CreateTime:    utils.TimePointerVal(resp.CreateTime),
		Status:        status,
		Total:         utils.Int64PointerVal(resp.Amount.Total),
		RefundAmount:  utils.Int64PointerVal(resp.Amount.Refund),
	}
}
