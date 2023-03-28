package file_op

import (
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/lfxnxf/emo-frame/logging"
	"github.com/lfxnxf/emo-server/conf"
	"go.uber.org/zap"
	"net/http"
	"os"
	"path"
)

type CecUpload struct {
	uploader *s3manager.Uploader
	c        *conf.Config
}

func NewCecUpload(c *conf.Config) FileOp {
	sess, err := session.NewSession(&aws.Config{
		Credentials:      credentials.NewStaticCredentials(c.Cec.AccessKeyID, c.Cec.AccessKeySecret, ""),
		Endpoint:         aws.String(c.CecOssConf.Endpoint),
		Region:           aws.String("wuhan-2"),
		DisableSSL:       aws.Bool(true),
		S3ForcePathStyle: aws.Bool(true),
	})
	if err != nil {
		panic(err)
	}

	uploader := s3manager.NewUploader(sess)

	return &CecUpload{
		uploader: uploader,
		c:        c,
	}
}

// UploadFile 上传本地文件
func (f *CecUpload) UploadFile(ctx context.Context, dir, file string) (string, error) {
	log := logging.For(ctx, "func", "UploadFile",
		zap.String("dir", dir),
		zap.String("file", file),
	)

	filename := path.Base(file)

	var obj string
	if len(dir) > 0 {
		obj = fmt.Sprintf("%s/%s", dir, filename)
	} else {
		obj = filename
	}

	err := f.Upload(ctx, filename, obj)
	if err != nil {
		log.Errorw("f.Upload error", zap.Error(err))
		return "", err
	}
	return fmt.Sprintf("%s/%s/%s", f.c.CecImageUrl, f.c.CecOssConf.BucketName, obj), nil
}

// UploadFileByStr 根据字符串生成图片并上传  //TODO  attachmentFlag bool
func (f *CecUpload) UploadFileByStr(ctx context.Context, dir, str string) (string, error) {
	log := logging.For(ctx, "func", "CreateFileByString",
		zap.String("dir", dir),
	)

	// 本地文件
	file, err := CreateFileByString(ctx, str, f.c.ImageDir)
	if err != nil {
		log.Errorw("CreateFileByString error", zap.Error(err))
		return "", err
	}

	filename := path.Base(file)

	// 远程oss地址
	obj := fmt.Sprintf("%s/%s", dir, filename)
	err = f.Upload(ctx, file, obj)
	if err != nil {
		log.Errorw("f.Upload error", zap.Error(err))
		return "", err
	}
	return fmt.Sprintf("%s/%s/%s", f.c.CecImageUrl, f.c.CecOssConf.BucketName, obj), nil
}

// UploadFileByRequest 客户端上传图片
func (f *CecUpload) UploadFileByRequest(ctx context.Context, request *http.Request, dir, formName string) (string, error) {
	log := logging.For(ctx, "func", "UploadFileByRequest",
		zap.String("dir", dir),
		zap.String("form_name", formName),
	)

	// 本地文件
	file, err := CreateFileByRequest(ctx, request, f.c.ImageDir, formName)
	if err != nil {
		log.Errorw("CreateFileByString error", zap.Error(err))
		return "", err
	}

	filename := path.Base(file)

	// 远程oss地址
	obj := fmt.Sprintf("%s/%s", dir, filename)
	err = f.Upload(ctx, file, obj)
	if err != nil {
		log.Errorw("f.Upload error", zap.Error(err))
		return "", err
	}
	return fmt.Sprintf("%s/%s/%s", f.c.CecImageUrl, f.c.CecOssConf.BucketName, obj), nil
}

func (f *CecUpload) Upload(ctx context.Context, filename, target string) error {
	log := logging.For(ctx, "func", "Upload",
		zap.String("filename", filename),
		zap.String("target", target),
	)
	file, err := os.Open(filename)
	if err != nil {
		log.Errorw("os.Open error", zap.Error(err))
		return err
	}

	_, err = f.uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String(f.c.CecOssConf.BucketName),
		Key:    aws.String(target),
		Body:   file,
	})
	if err != nil {
		log.Errorw("f.uploader.Upload error", zap.Error(err))
		return err
	}
	return nil
}
