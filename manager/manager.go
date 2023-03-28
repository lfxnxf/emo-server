package manager

import (
	"context"
	"github.com/lfxnxf/emo-frame/inits/proxy"
	"github.com/lfxnxf/emo-frame/tools/syncx"
	"github.com/lfxnxf/emo-server/conf"
	"github.com/lfxnxf/emo-server/dao"
	"github.com/lfxnxf/emo-server/tools/file_op"
	"github.com/lfxnxf/emo-server/tools/qrcode"
	"github.com/lfxnxf/emo-server/tools/sms"
	"github.com/lfxnxf/emo-server/tools/wx_pay"
	"sync"
)

var mqMap = map[string]string{}

type MqManger struct {
	Topic    string `json:"topic"`
	Consumer *proxy.KafkaConsumer
}

// Manager represents middleware component
// such as, kafka, http client or rpc client, etc.
type Manager struct {
	c            *conf.Config
	dao          *dao.Dao
	lock         sync.Mutex
	singleFlight syncx.SingleFlight
	wxPay        *wx_pay.WxPay
	kafkaPro     *proxy.KafkaSyncProducer
	qrcode       *qrcode.Qrcode
	fileOp       file_op.FileOp
	MqMap        map[string]MqManger
	sms          *sms.Sms
}

func New(conf *conf.Config) *Manager {
	m := &Manager{
		c:            conf,
		dao:          dao.New(conf),
		singleFlight: syncx.NewSingleFlight(),
		wxPay:        wx_pay.New(conf.WechatProgram),
		kafkaPro:     proxy.InitKafkaSyncProducer("kafka_pro"),
		qrcode:       qrcode.NewQrcode(conf),
		fileOp:       file_op.NewFileOp(conf.ImageTag, conf),
		sms:          sms.New(conf),
	}
	m.initMqMap()
	return m
}

func (m *Manager) initMqMap() {
	m.MqMap = make(map[string]MqManger)
	m.lock.Lock()
	for k, v := range mqMap {
		m.MqMap[k] = MqManger{
			Topic:    v,
			Consumer: proxy.InitKafkaConsumer(k),
		}
	}
	m.lock.Unlock()
}

func (m *Manager) Ping(ctx context.Context) error {
	return nil
}

func (m *Manager) Close() error {
	return nil
}

type GetVisToken func(context.Context, string, string) (string, error)
