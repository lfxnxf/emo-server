package file_op

import (
	"context"
	"fmt"
	"github.com/lfxnxf/emo-frame/inits"
	"github.com/lfxnxf/emo-server/conf"
	"os"
	"testing"
)

var cfg *conf.Config

func InitTest() {
	inits.Init(
		inits.ConfigPath("../../app/config/local/config.toml"),
		inits.Once(),
		inits.ConfigInstance(new(conf.Config)),
	)

	// init local config
	cfg = conf.Init()

	// create images dir
	if len(cfg.ImageDir) > 0 {
		_ = os.MkdirAll(cfg.ImageDir, 0777)
	}
}

func TestUpload(t *testing.T) {
	InitTest()
	fileOp := NewFileOp(Cec, cfg)
	res, err := fileOp.UploadFile(context.Background(), "qrcode", "./a.txt")
	if err != nil {
		panic(err)
	}
	fmt.Println(res)
}
