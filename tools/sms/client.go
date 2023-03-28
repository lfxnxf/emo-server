package sms

import (
	"context"
	openApi "github.com/alibabacloud-go/darabonba-openapi/client"
	smsApi "github.com/alibabacloud-go/dysmsapi-20170525/v2/client"
	util "github.com/alibabacloud-go/tea-utils/service"
	"github.com/alibabacloud-go/tea/tea"
	"github.com/lfxnxf/emo-frame/logging"
	"github.com/lfxnxf/emo-frame/zd_error"
	"github.com/lfxnxf/emo-server/conf"
	"go.uber.org/zap"
)

type Sms struct {
	client *smsApi.Client
}

func New(c *conf.Config) *Sms {
	config := &openApi.Config{
		AccessKeyId:     tea.String(c.Ali.AccessKeyID),
		AccessKeySecret: tea.String(c.Ali.AccessKeySecret),
	}
	config.Endpoint = tea.String(endpoint)
	client, err := smsApi.NewClient(config)
	if err != nil {
		logging.Fatal(err)
	}
	return &Sms{client: client}
}

func (s *Sms) SendMsg(ctx context.Context, phone, sign, template, param string) error {
	log := logging.For(ctx, "func", "SendMsg",
		zap.String("phone", phone),
		zap.String("sign", sign),
		zap.String("template", template),
		zap.String("param", param),
	)

	sendSmsRequest := &smsApi.SendSmsRequest{
		PhoneNumbers:  tea.String(phone),
		SignName:      tea.String(sign),
		TemplateCode:  tea.String(template),
		TemplateParam: tea.String(param),
	}
	runtime := &util.RuntimeOptions{}

	resp, err := s.client.SendSmsWithOptions(sendSmsRequest, runtime)
	if err != nil {
		log.Errorw("s.client.SendSmsWithOptions error", zap.Error(err))
		return err
	}

	log.Infow("send msg", zap.Any("resp", resp.Body))

	if resp.Body.Code != nil && *resp.Body.Code != "OK" {
		return zd_error.AddSpecialError(500, *resp.Body.Message)
	}

	return nil
}
