package service

import (
	"context"
	"github.com/lfxnxf/emo-frame/tools/syncx"
	"github.com/lfxnxf/emo-server/conf"
	"github.com/lfxnxf/emo-server/dao"
	"github.com/lfxnxf/emo-server/manager"
	"github.com/robfig/cron"
)

const ()

const (
	defaultPage  = 1
	defaultLimit = 20

	// redis zset 默认分页数据
	zsetDefaultStart = 0
	zsetDefaultStop  = 19
	zsetDefaultLimit = 20
)

type Service struct {
	c *conf.Config

	// dao: db handler
	dao *dao.Dao

	// manager: other client(s), other middleware(s)
	mgr *manager.Manager

	singleFlight syncx.SingleFlight
}

func New(c *conf.Config) *Service {
	return &Service{
		c:            c,
		dao:          dao.New(c),
		mgr:          manager.New(c),
		singleFlight: syncx.NewSingleFlight(),
	}
}

// Ping check service's resource status
func (s *Service) Ping(ctx context.Context) error {
	return s.dao.Ping(ctx)
}

// StartConsume 启动kafka消费
func (s *Service) StartConsume(ctx context.Context) {

}

func (s *Service) StartCron() {
	c := cron.New()

	c.Start()
}

func (s *Service) genPageLimit(page, limit int64) (int64, int64) {
	if page == 0 {
		page = defaultPage
	}
	if limit == 0 {
		limit = defaultLimit
	}
	return page, limit
}
