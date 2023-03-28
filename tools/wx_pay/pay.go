package wx_pay

import (
	"context"
	"github.com/lfxnxf/emo-frame/logging"
	"github.com/lfxnxf/emo-server/conf"
	"github.com/wechatpay-apiv3/wechatpay-go/core"
	"github.com/wechatpay-apiv3/wechatpay-go/core/option"
	"github.com/wechatpay-apiv3/wechatpay-go/utils"
	"go.uber.org/zap"
)

type WxPay struct {
	client *core.Client
	config conf.WechatProgramConf
}

func New(config conf.WechatProgramConf) *WxPay {
	// 使用 local_utils 提供的函数从本地文件中加载商户私钥，商户私钥会用来生成请求的签名
	mchPrivateKey, err := utils.LoadPrivateKeyWithPath(config.PrivateKeyPath)
	if err != nil {
		logging.Error("load merchant private key error", zap.Error(err))
		return &WxPay{}
	}
	// 使用商户私钥等初始化 client，并使它具有自动定时获取微信支付平台证书的能力
	opts := []core.ClientOption{
		option.WithWechatPayAutoAuthCipher(config.MchID, config.MchCertificateSerialNumber, mchPrivateKey, config.MchAPIv3Key),
	}
	client, err := core.NewClient(context.Background(), opts...)
	if err != nil {
		logging.Errorw("new wechat pay client err:%s", zap.Error(err))
		return &WxPay{}
	}
	return &WxPay{
		client: client,
		config: config,
	}
}
