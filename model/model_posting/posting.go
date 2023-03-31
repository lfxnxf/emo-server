package model_posting

import "time"

const (
	_      = iota
	ScoreS //S
	ScoreA //A
	ScoreB //B
	ScoreC //C
)

const (
	PostingTypeNormal   = 1 // 普通
	PostingTypeBoutique = 2 // 精选
)

const (
	TablePosting = "posting"
)

// Posting 帖子表
type Posting struct {
	ID              int64     `gorm:"column:id" json:"id"`                                 //自增id
	Uid             int64     `gorm:"column:uid" json:"uid"`                               //用户id
	Content         string    `gorm:"column:content" json:"content"`                       //帖子内容
	Images          string    `gorm:"column:images" json:"images"`                         //图片
	Score           int64     `gorm:"column:score" json:"score"`                           //质量分，1:S,2:A,3:B,4:C
	UserType        int64     `gorm:"column:user_type" json:"user_type"`                   //1：自然贴，2：马甲贴
	PostingType     int64     `gorm:"column:posting_type" json:"posting_type"`             //类型，1：普通，2：精选
	AuditStatus     int64     `gorm:"column:audit_status" json:"audit_status"`             //审核状态，1：未审核，2：审核成功，10：审核未通过
	AuditFailReason string    `gorm:"column:audit_fail_reason" json:"audit_fail_reason"`   //审核失败原因
	Status          int64     `gorm:"column:status" json:"status"`                         //状态，1：未发布，2:已发布，101:已删除
	CreateTime      time.Time `gorm:"column:create_time; default:null" json:"create_time"` //创建时间
}

func (p *Posting) TableName() string {
	return TablePosting
}

func (p *Posting) ScoreText() string {
	scoreMap := map[int64]string{
		ScoreS: "S",
		ScoreA: "A",
		ScoreB: "B",
		ScoreC: "C",
	}
	return scoreMap[p.Score]
}

func (p *Posting) TypeText() string {
	typeMap := map[int64]string{
		PostingTypeNormal:   "普通",
		PostingTypeBoutique: "精选",
	}
	return typeMap[p.PostingType]
}

type SearchPosting struct {
	Posting
	Nickname        string `json:"nickname" gorm:"column:nickname"`
	Subjects        string `json:"subjects" gorm:"column:subjects"`
	LikeNum         int64  `json:"like_num" gorm:"column:like_num"`
	HumanCommentNum int64  `json:"human_comment_num" gorm:"column:human_comment_num"`
	AllCommentNum   int64  `json:"all_comment_num" gorm:"column:all_comment_num"`
}
