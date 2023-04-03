package model_posting

import "github.com/lfxnxf/emo-server/model"

type AddPostingReq struct {
	Uid        int64    `json:"uid"`
	SubjectIds []int64  `json:"subject_ids"`
	Content    string   `json:"content"`
	Images     []string `json:"images"`
}

type GetAllSubjectReply struct {
	List []Subject `json:"list"`
}

type SearchPostingReq struct {
	PostingId   int64           `json:"posting_id" form:"posting_id"`
	Uid         int64           `json:"uid" form:"uid"`
	PostingType int64           `json:"posting_type" form:"posting_type"`
	UserType    int64           `json:"user_type" form:"user_type"`
	Score       int64           `json:"score" form:"score"`
	StartAt     string          `json:"start_at" form:"start_at"`
	EndAt       string          `json:"end_at" form:"end_at"`
	Content     string          `json:"content" form:"content"`
	Page        int64           `json:"page" form:"page"`
	Limit       int64           `json:"limit" form:"limit"`
	OrderBy     []model.OrderBy `json:"order_by" form:"order_by"`
}

type SearchPostingReply struct {
	Page  int64              `json:"page"`
	Limit int64              `json:"limit"`
	Total int64              `json:"total"`
	List  []AdminPostingInfo `json:"list"`
}

type AdminPostingInfo struct {
	PostingId       int64    `json:"posting_id"`
	Uid             int64    `json:"uid"`
	Content         string   `json:"content"`
	Images          []string `json:"images"`
	Subjects        []string `json:"subjects"`
	Score           int64    `json:"score"`
	ScoreText       string   `json:"score_text"`
	LikeNum         int64    `json:"like_num"`
	HumanCommentNum int64    `json:"human_comment_num"`
	AllCommentNum   int64    `json:"all_comment_num"`
	CreateTime      string   `json:"create_time"`
	PostingType     int64    `json:"posting_type"`
	PostingTypeText string   `json:"posting_type_text"`
}

const (
	StatusRefine       = 1
	StatusCancelRefine = 2
)

type RefineOrCancelPostingReq struct {
	PostingId int64 `json:"posting_id"`
	Status    int64 `json:"status"`
}
