package manager

import (
	"context"
	"github.com/lfxnxf/emo-frame/logging"
	"go.uber.org/zap"
	"net/http"
)

func (m *Manager) UploadFile(ctx context.Context, request *http.Request, dir, formName string) (string, error) {
	log := logging.For(ctx, "func", "UploadFile")
	file, err := m.fileOp.UploadFileByRequest(ctx, request, dir, formName)
	if err != nil {
		log.Errorw("m.fileOp.UploadFileByRequest error", zap.Error(err))
		return "", err
	}
	log.Infow("success!")
	return file, nil
}

func (m *Manager) UploadFileByFile(ctx context.Context, dir, file string) (string, error) {
	log := logging.For(ctx, "func", "UploadFile")
	file, err := m.fileOp.UploadFile(ctx, dir, file)
	if err != nil {
		log.Errorw("m.fileOp.UploadFile error", zap.Error(err))
		return "", err
	}
	log.Infow("success!")
	return file, nil
}

func (m *Manager) UploadFileByStr(ctx context.Context, dir, file string) (string, error) {
	log := logging.For(ctx, "func", "UploadFile")
	// 上传到oss
	url, err := m.fileOp.UploadFileByStr(ctx, dir, file)
	if err != nil {
		log.Errorw("m.fileOp.UploadFileByStr", zap.Error(err))
		return "", err
	}
	return url, nil
}
