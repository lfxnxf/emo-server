package model_posting

import "time"

const (
	TablePostingStatistics = "posting_statistics"
)

const (
	_                               = iota
	BusinessTypePosting             // 帖子
	BusinessTypePostingComment      // 帖子评论
	BusinessTypePostingCommentReply // 帖子评论回复
)

const (
	_                    = iota
	TypeNormalThumbNum   // 自然人点赞数量
	TypeRobotThumbNum    // 马甲人点赞数量
	TypeNormalCommentNum // 自然人评论数量
	TypeRobotCommentNum  // 马甲人评论数量
)

// PostingStatistics  帖子相关统计表
type PostingStatistics struct {
	ID             int64     `gorm:"column:id" json:"id"`                                 //  自增id
	BusinessType   int64     `gorm:"column:business_type" json:"business_type"`           //  点赞类型，1：帖子，2：帖子评论，3：评论回复
	BusinessId     int64     `gorm:"column:business_id" json:"business_id"`               //  根据类型获取不同id
	StatisticsType int64     `gorm:"column:statistics_type" json:"statistics_type"`       //  统计类型，1：自然人点赞数量，2：全部点赞数量，3：自然人评论数量，4：全部评论数量
	Num            int64     `gorm:"column:num" json:"num"`                               //  数量
	CreateTime     time.Time `gorm:"column:create_time; default:null" json:"create_time"` //  创建时间
}

func (PostingStatistics) TableName() string {
	return TablePostingStatistics
}
