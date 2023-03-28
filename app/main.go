package main

import (
	"context"
	"flag"
	"github.com/lfxnxf/emo-frame/inits"
	"github.com/lfxnxf/emo-server/conf"
	"github.com/lfxnxf/emo-server/server/http"
	"github.com/lfxnxf/emo-server/service"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func init() {
	configS := flag.String("config", "./app/config/local/config.toml", "Configuration file")
	flag.Parse()

	inits.Init(
		inits.ConfigPath(*configS),
		inits.Once(),
		inits.ConfigInstance(new(conf.Config)),
	)
}

func main() {
	// init local config
	cfg := conf.Init()

	// create images dir
	if len(cfg.ImageDir) > 0 {
		_ = os.MkdirAll(cfg.ImageDir, 0777)
	}

	// create a service instance
	srv := service.New(cfg)

	// 开启消费
	go srv.StartConsume(context.Background())

	// 开启定时任务
	go srv.StartCron()

	// init and start http server
	http.Init(srv)

	defer http.Shutdown()

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT)
	for {
		s := <-sigChan
		log.Printf("get a signal %s\n", s.String())
		switch s {
		case syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT:
			log.Println("api server exit now...")
			return
		case syscall.SIGHUP:
		default:
		}
	}
}
