package model_like

import "time"

const (
	TableLikeRecord = "like_record"
)
const (
	_                           = iota
	LikeTypePosting             // 帖子
	LikeTypePostingComment      // 帖子评论
	LikeTypePostingCommentReply // 帖子评论回复
)

const (
	LikeRecordStatusThumb    = 1   // 已点赞
	LikeRecordStatusCanceled = 101 // 已取消
)

// LikeRecord  点赞记录表
type LikeRecord struct {
	ID           int64     `gorm:"column:id" json:"id"`                                 // 自增id
	Uid          int64     `gorm:"column:uid" json:"uid"`                               // 用户id
	BusinessType int64     `gorm:"column:business_type" json:"business_type"`           // 业务类型，1：帖子，2：帖子评论，3：评论回复
	BusinessId   int64     `gorm:"column:business_id" json:"business_id"`               // 根据类型获取不同id
	Status       int64     `gorm:"column:status" json:"status"`                         // 状态，1:已点赞，101:已取消
	CreateTime   time.Time `gorm:"column:create_time; default:null" json:"create_time"` // 创建时间
}

func (LikeRecord) TableName() string {
	return TableLikeRecord
}
