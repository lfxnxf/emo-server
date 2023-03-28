package dao

import (
	"context"
	"github.com/jinzhu/gorm"
	"github.com/lfxnxf/emo-frame/inits/proxy"
	"github.com/lfxnxf/emo-frame/resource/redis"
	"github.com/lfxnxf/emo-server/conf"
	"sync"
)

const (
	RedisClient       = "api.redis"
	WechatRedisClient = "wechat.redis"

	RankingListZsetMaxValue = 1000 // 排行榜中最多只存1000个数据
)

// Dao represents data access object
type Dao struct {
	c           *conf.Config
	cache       *proxy.Redis
	wechatCache *proxy.Redis
	db          *proxy.SQL
	scriptMap   *sync.Map
}

func New(c *conf.Config) *Dao {
	d := &Dao{
		c:           c,
		cache:       proxy.InitRedis(RedisClient),
		db:          proxy.InitSQL("api.db"),
		wechatCache: proxy.InitRedis(WechatRedisClient),
	}
	d.InitScript()
	return d
}

func (d *Dao) InitScript() {
	// 加载lua
	scriptMap := new(sync.Map)

	d.scriptMap = scriptMap
}

func (d *Dao) Begin() *gorm.DB {
	return d.db.Master().Begin()
}

func (d *Dao) GetCacheClient() *redis.Redis {
	return d.cache.Redis
}

// Ping check db resource status
func (d *Dao) Ping(ctx context.Context) error {
	return nil
}

// Close release resource
func (d *Dao) Close() error {
	return nil
}

func (d *Dao) Tx(tx *gorm.DB) *gorm.DB {
	if tx == nil {
		tx = d.db.Master().DB
	}
	return tx
}
