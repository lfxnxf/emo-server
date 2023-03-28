package manager

import (
	"context"
	"errors"
	"github.com/lfxnxf/emo-frame/logging"
	"go.uber.org/zap"
	"math/rand"
	"strconv"
)

func (m *Manager) SyncProducerByMap(ctx context.Context, key, msg string) (partition int32, offset int64, err error) {
	log := logging.For(ctx, "func", "SyncProducerByMap",
		zap.Any("msg", msg),
	)

	var (
		mq MqManger
		ok bool
	)
	if mq, ok = m.MqMap[key]; !ok {
		return 0, 0, errors.New("key not exist")
	}

	partitionKey := strconv.FormatInt(rand.Int63(), 10)
	partition, offset, err = m.kafkaPro.SendSyncMsg(ctx, mq.Topic, partitionKey, []byte(msg))

	if err != nil {
		log.Errorw("SendSyncMsg error", zap.Error(err))
		return 0, 0, err
	}
	log.Infow("success")
	return
}
