package service

import (
	"context"
	"github.com/lfxnxf/emo-frame/logging"
	"github.com/lfxnxf/emo-frame/utils"
	"github.com/lfxnxf/emo-server/model"
	"github.com/lfxnxf/emo-server/model/model_posting"
	"go.uber.org/zap"
	"strings"
)

// AddPosting 新增帖子
func (s *Service) AddPosting(ctx context.Context, req model_posting.AddPostingReq) (interface{}, error) {
	log := logging.For(ctx, "func", "AddPosting",
		zap.Any("req", req),
	)

	_, err := s.dao.InsertPosting(ctx, nil, model_posting.Posting{
		Uid:         req.Uid,
		Content:     req.Content,
		Images:      strings.Join(req.Images, ","),
		UserType:    model.UserTypeRobot,
		PostingType: model_posting.PostingTypeNormal,
		AuditStatus: model_posting.AuditStatusPass,
		Status:      model_posting.StatusPublished,
	})
	if err != nil {
		log.Errorw("s.dao.InsertPosting error", zap.Error(err))
		return nil, err
	}

	log.Infow("success!")
	return nil, nil
}

// GetAllSubject 获取所有话题
func (s *Service) GetAllSubject(ctx context.Context) (interface{}, error) {
	log := logging.For(ctx, "func", "GetAllSubject")

	subject, err := s.dao.GetAllSubject(ctx)
	if err != nil {
		log.Errorw("s.dao.GetAllSubject error", zap.Error(err))
		return nil, err
	}

	log.Infow("success!")
	return model_posting.GetAllSubjectReply{List: subject}, nil
}

func (s *Service) SearchPosting(ctx context.Context, req model_posting.SearchPostingReq) (interface{}, error) {
	log := logging.For(ctx, "func", "SearchPosting",
		zap.Any("req", req),
	)
	page, limit := s.genPageLimit(req.Page, req.Limit)

	resp := model_posting.SearchPostingReply{
		Page:  page,
		Limit: limit,
		List:  make([]model_posting.AdminPostingInfo, 0),
	}

	// 组装查询参数
	where := make(map[string]interface{})

	if req.PostingId > 0 {
		where["id"] = req.PostingId
	}

	if req.Uid > 0 {
		where["uid"] = req.PostingId
	}

	if len(req.Content) > 0 {
		where["content"] = req.Content
	}

	if req.UserType > 0 {
		where["user_type"] = req.UserType
	}

	if req.PostingType > 0 {
		where["posting_type"] = req.PostingType
	}

	if len(req.StartAt) > 0 {
		where["start_at"] = req.StartAt
	}

	if len(req.EndAt) > 0 {
		where["end_at"] = req.EndAt
	}

	list, total, err := s.dao.SearchPosting(ctx, where, req.OrderBy, page, limit, true)
	if err != nil {
		log.Errorw("s.dao.SearchPosting error", zap.Error(err))
		return nil, err
	}

	resp.Total = total

	//组装数据
	for _, v := range list {
		info := model_posting.AdminPostingInfo{
			PostingId:       v.ID,
			Uid:             v.Uid,
			Content:         v.Content[0:300],
			Images:          strings.Split(v.Images, ","),
			Subjects:        strings.Split(v.Subjects, ","),
			Score:           v.Score,
			ScoreText:       v.ScoreText(),
			LikeNum:         v.LikeNum,
			HumanCommentNum: v.HumanCommentNum,
			AllCommentNum:   v.AllCommentNum,
			CreateTime:      v.CreateTime.Format(utils.TimeFormatYYYYMMDDHHmmSS),
			PostingType:     v.PostingType,
			PostingTypeText: v.TypeText(),
		}
		resp.List = append(resp.List, info)
	}

	log.Infow("success!")
	return resp, nil
}
