package manager

import (
	"github.com/lfxnxf/emo-frame/inits"
	"github.com/lfxnxf/emo-server/conf"
	"github.com/lfxnxf/emo-server/dao"
	"os"
)

var m *Manager
var d *dao.Dao

func InitTest() {
	inits.Init(
		inits.ConfigPath("../app/config/local/config.toml"),
		inits.Once(),
		inits.ConfigInstance(new(conf.Config)),
	)

	// init local config
	cfg := conf.Init()

	// create images dir
	if len(cfg.ImageDir) > 0 {
		_ = os.MkdirAll(cfg.ImageDir, 0777)
	}

	// create a service instance
	m = New(cfg)
	d = dao.New(cfg)
}
