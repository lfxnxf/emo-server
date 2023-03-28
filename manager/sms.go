package manager

import (
	"context"
	jsoniter "github.com/json-iterator/go"
	"github.com/lfxnxf/emo-frame/logging"
	"go.uber.org/zap"
)

// 发送短信
func (m *Manager) SendSms(ctx context.Context, phone, sign, template string, params map[string]interface{}) error {
	log := logging.For(ctx, "func", "SendSms",
		zap.String("phone", phone),
		zap.String("sign", sign),
		zap.String("template", template),
		zap.Any("params", params),
	)

	marshal, err := jsoniter.Marshal(params)
	if err != nil {
		log.Errorw("jsoniter marshal", zap.Error(err))
		return err
	}

	// 发送短信
	err = m.sms.SendMsg(ctx, phone, sign, template, string(marshal))
	if err != nil {
		log.Errorw("m.sms.SendMsg error", zap.Error(err))
		return err
	}

	log.Infow("success!")
	return nil
}
