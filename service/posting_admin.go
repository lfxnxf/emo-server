package service

import (
	"context"
	"github.com/lfxnxf/emo-frame/logging"
	"github.com/lfxnxf/emo-server/model/model_posting"
	"go.uber.org/zap"
)

func (s *Service) AddPosting(ctx context.Context, req model_posting.AddPostingReq) (interface{}, error) {
	log := logging.For(ctx, "func", "AddPosting",
		zap.Any("req", req),
	)



	log.Infow("success!")
	return nil, nil
}

