package qrcode

import (
	"context"
	"fmt"
	"github.com/lfxnxf/emo-frame/logging"
	"github.com/lfxnxf/emo-frame/utils"
	"github.com/lfxnxf/emo-server/conf"
	"github.com/skip2/go-qrcode"
	"go.uber.org/zap"
	"time"
)

type Qrcode struct {
	c *conf.Config
}

func NewQrcode(c *conf.Config) *Qrcode {
	return &Qrcode{c: c}
}

func (q *Qrcode) CreateQrcode(ctx context.Context, content string, size int) (string, error) {
	log := logging.For(ctx, "func", "CreateQrcode",
		zap.String("content", content),
		zap.Int("size", size),
	)

	// 生成二维码
	qrcodeName := utils.Md5(fmt.Sprintf("%d%s", time.Now().Nanosecond(), utils.GenRandString(10, 3)))

	file := fmt.Sprintf("%s/%s.png", q.c.ImageDir, qrcodeName)
	err := qrcode.WriteFile(content, qrcode.Medium, size, file)
	if err != nil {
		log.Errorw("qrcode.WriteFile error", zap.Error(err))
		return "", err
	}
	return file, nil
}
