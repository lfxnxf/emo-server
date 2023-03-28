package file_op

import (
	"context"
	"github.com/lfxnxf/emo-server/conf"
	"net/http"
)

const (
	Ali = "ali"
	Cec = "cec"
)

type FileOp interface {
	UploadFile(ctx context.Context, dir, file string) (string, error)
	UploadFileByStr(ctx context.Context, dir, str string) (string, error)
	UploadFileByRequest(ctx context.Context, request *http.Request, dir, formName string) (string, error)
}

func NewFileOp(key string, c *conf.Config) FileOp {
	return map[string]FileOp{
		Ali: NewAliUpload(c),
		Cec: NewCecUpload(c),
	}[key]
}
