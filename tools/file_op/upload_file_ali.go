package file_op

import (
	"context"
	"fmt"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/lfxnxf/emo-frame/logging"
	"github.com/lfxnxf/emo-server/conf"
	"go.uber.org/zap"
	"net/http"
	"path"
)

type AliUpload struct {
	Oss *Oss
	c   *conf.Config
}

type Oss struct {
	client *oss.Client
	bucket *oss.Bucket
}

func NewAliUpload(c *conf.Config) FileOp {
	client, err := oss.New(
		c.AliOssConf.Endpoint,
		c.Ali.AccessKeyID,
		c.Ali.AccessKeySecret,
		oss.Timeout(10, 120),
	)
	if err != nil {
		panic(err)
	}

	bucket, err := client.Bucket(c.AliOssConf.BucketName)
	if err != nil {
		panic(err)
	}

	return &AliUpload{
		c: c,
		Oss: &Oss{
			client: client,
			bucket: bucket,
		},
	}
}

// UploadFile 上传本地文件
func (f *AliUpload) UploadFile(ctx context.Context, dir, file string) (string, error) {
	log := logging.For(ctx, "func", "UploadFile",
		zap.String("dir", dir),
		zap.String("file", file),
	)

	filename := path.Base(file)

	// 远程oss地址
	obj := fmt.Sprintf("%s/%s", dir, filename)
	err := f.Oss.bucket.PutObjectFromFile(obj, file)
	if err != nil {
		log.Errorw("f.Oss.bucket.PutObjectFromFile error", zap.Error(err))
		return "", err
	}
	return fmt.Sprintf("%s/%s", f.c.ImageUrl, obj), nil
}

// 根据字符串生成图片并上传
func (f *AliUpload) UploadFileByStr(ctx context.Context, dir, str string) (string, error) {
	log := logging.For(ctx, "func", "CreateFileByString",
		zap.String("dir", dir),
	)

	// 本地文件
	file, err := CreateFileByString(ctx, str, f.c.ImageDir)
	if err != nil {
		log.Errorw("f.CreateFileByString error", zap.Error(err))
		return "", err
	}

	filename := path.Base(file)

	// 远程oss地址
	obj := fmt.Sprintf("%s/%s", dir, filename)
	err = f.Oss.bucket.PutObjectFromFile(obj, file)
	if err != nil {
		log.Errorw("f.Oss.bucket.PutObjectFromFile error", zap.Error(err))
		return "", err
	}
	return fmt.Sprintf("%s/%s", f.c.ImageUrl, obj), nil
}

// UploadFileByRequest 客户端上传图片
func (f *AliUpload) UploadFileByRequest(ctx context.Context, request *http.Request, dir, formName string) (string, error) {
	log := logging.For(ctx, "func", "UploadFileByRequest",
		zap.String("dir", dir),
		zap.String("form_name", formName),
	)

	// 本地文件
	file, err := CreateFileByRequest(ctx, request, f.c.ImageDir, formName)
	if err != nil {
		log.Errorw("f.CreateFileByString error", zap.Error(err))
		return "", err
	}

	filename := path.Base(file)

	// 远程oss地址
	obj := fmt.Sprintf("%s/%s", dir, filename)
	err = f.Oss.bucket.PutObjectFromFile(obj, file)
	if err != nil {
		log.Errorw("f.Oss.bucket.PutObjectFromFile error", zap.Error(err))
		return "", err
	}
	return fmt.Sprintf("%s/%s", f.c.ImageUrl, obj), nil
}
