package model_posting

import "time"

const (
	TablePostingComment = "posting_comment"
)

// PostingComment  评论表
type PostingComment struct {
	ID              int64     `gorm:"column:id" json:"id"`                                 // 自增id
	Uid             int64     `gorm:"column:uid" json:"uid"`                               // 用户id
	PostingId       int64     `gorm:"column:posting_id" json:"posting_id"`                 // 帖子id
	Attribute       int64     `gorm:"column:attribute" json:"attribute"`                   // 属性，1：自然人，2：马甲人
	Content         string    `gorm:"column:content" json:"content"`                       // 评论内容
	AuditStatus     int64     `gorm:"column:audit_status" json:"audit_status"`             // 审核状态，1：未审核，2：审核通过，10：审核未通过
	AuditFailReason string    `gorm:"column:audit_fail_reason" json:"audit_fail_reason"`   // 审核未通过原因
	Status          int64     `gorm:"column:status" json:"status"`                         // 状态，1：未发布，2:已发布，101:已删除
	CreateTime      time.Time `gorm:"column:create_time; default:null" json:"create_time"` // 创建时间
}

func (PostingComment) TableName() string {
	return TablePostingComment
}
