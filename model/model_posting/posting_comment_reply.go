package model_posting

import "time"

const (
	TablePostingCommentReply = "posting_comment_reply"
)

// PostingCommentReply  评论回复表
type PostingCommentReply struct {
	ID               int64     `gorm:"column:id" json:"id"`                                 //  自增id
	PostingId        int64     `gorm:"column:posting_id" json:"posting_id"`                 //  帖子id
	CommentId        int64     `gorm:"column:comment_id" json:"comment_id"`                 //  评论id
	ThreadStarterUid int64     `gorm:"column:thread_starter_uid" json:"thread_starter_uid"` //  楼主uid
	Sender           int64     `gorm:"column:sender" json:"sender"`                         //  用户id
	Receiver         int64     `gorm:"column:receiver" json:"receiver"`                     //  被回复用户id
	ReceiveReplyId   int64     `gorm:"column:receive_reply_id" json:"receive_reply_id"`     //  被回复内容id
	Attribute        int64     `gorm:"column:attribute" json:"attribute"`                   //  属性，1：自然人，2：马甲人
	Content          string    `gorm:"column:content" json:"content"`                       //  回复内容
	AuditStatus      int64     `gorm:"column:audit_status" json:"audit_status"`             //  审核状态，1：未审核，2：审核通过，10：审核未通过
	AuditFailReason  string    `gorm:"column:audit_fail_reason" json:"audit_fail_reason"`   //  审核未通过原因
	Status           int64     `gorm:"column:status" json:"status"`                         //  状态，1：未发布，2:已发布，101:已删除
	CreateTime       time.Time `gorm:"column:create_time; default:null" json:"create_time"` //  创建时间
}

func (PostingCommentReply) TableName() string {
	return TablePostingCommentReply
}
