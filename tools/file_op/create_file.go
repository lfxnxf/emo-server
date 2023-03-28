package file_op

import (
	"context"
	"fmt"
	"github.com/lfxnxf/emo-frame/logging"
	"github.com/lfxnxf/emo-frame/utils"
	"go.uber.org/zap"
	"io"
	"net/http"
	"os"
	"time"
)

func CreateFileByString(ctx context.Context, str, imageDir string) (string, error) {
	log := logging.For(ctx, "func", "CreateFileByString")

	fileName := utils.Md5(fmt.Sprintf("%d%s", time.Now().Nanosecond(), utils.GenRandString(10, 3)))

	// 生成二维码
	file := fmt.Sprintf("%s/%s.png", imageDir, fileName)
	content, err := os.Create(file)
	defer func() {
		_ = content.Close()
	}()
	if err != nil {
		log.Errorw("os.Create error", zap.Error(err))
		return "", err
	}

	_, err = io.WriteString(content, str)
	if err != nil {
		log.Errorw("io.WriteString error", zap.Error(err))
		return "", err
	}
	return file, nil
}

func CreateFileByRequest(ctx context.Context, request *http.Request, imageDir, formName string) (string, error) {
	log := logging.For(ctx, "func", "CreateFileByRequest")

	uploadFile, _, err := request.FormFile(formName)
	if err != nil {
		log.Errorw("request.FormFile error", zap.Error(err))
		return "", err
	}

	defer func() {
		_ = uploadFile.Close()
	}()

	file := fmt.Sprintf("%s/%s.png", imageDir, utils.Md5(fmt.Sprintf("%d%s", time.Now().Nanosecond(), utils.GenRandString(10, 3))))
	saveFile, err := os.OpenFile(file, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		log.Errorw("os.OpenFile error", zap.Error(err))
		return "", err
	}

	defer func() {
		_ = saveFile.Close()
	}()

	_, _ = io.Copy(saveFile, uploadFile)

	return file, nil
}
