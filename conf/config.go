package conf

import (
	"github.com/lfxnxf/emo-frame/config"
	"github.com/lfxnxf/emo-frame/inits"
)

type Config struct {
	config.Config
	Env           string            `toml:"env"`
	WechatProgram WechatProgramConf `toml:"wechat_program"`
	ImageDir      string            `toml:"image_dir"`
	ImageUrl      string            `toml:"image_url"`
	CecImageUrl   string            `toml:"cec_image_url"`
	Ali           Access            `toml:"ali"`
	AliOssConf    OssConf           `toml:"ali_oss_conf"`
	Cec           Access            `toml:"cec"`
	CecOssConf    OssConf           `toml:"cec_oss_conf"`
	ImageTag      string            `toml:"image_tag"`
}

type WechatProgramConf struct {
	AppId                      string `toml:"app_id"`
	AppSecret                  string `toml:"app_secret"`
	MchID                      string `toml:"mch_id"`
	MchCertificateSerialNumber string `toml:"mch_certificate_serial_number"`
	MchAPIv3Key                string `toml:"mch_ap_iv_3_key"`
	PrivateKeyPath             string `toml:"private_key_path"`
	PayNotifyUrl               string `toml:"pay_notify_url"`
	RefundNotifyUrl            string `toml:"refund_notify_url"`
}

type Access struct {
	AccessKeyID     string `toml:"access_key_id"`
	AccessKeySecret string `toml:"access_key_secret"`
}

type OssConf struct {
	Endpoint   string `toml:"endpoint"`
	BucketName string `toml:"bucket_name"`
}

func (c *Config) GetBase() config.Config {
	return c.Config
}

func Init() *Config {
	c, ok := inits.GetConfigInstance().(*Config)
	if !ok {
		panic("instance is not Config")
	}
	return c
}
