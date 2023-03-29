package model_posting

import "time"

const (
	TablePostingSubject = "posting_subject"

	PostingSubjectStatusNormal  = 1
	PostingSubjectStatusDeleted = 101
)

// PostingSubject 话题帖子关联表
type PostingSubject struct {
	ID         int64     `gorm:"column:id" json:"id"`                                 //自增id
	PostingId  int64     `gorm:"column:posting_id" json:"posting_id"`                 //帖子id
	SubjectId  int64     `gorm:"column:subject_id" json:"subject_id"`                 //话题id
	Status     int64     `gorm:"column:status" json:"status"`                         //状态，1:正常，101:删除
	CreateTime time.Time `gorm:"column:create_time; default:null" json:"create_time"` //创建时间
}

func (p *PostingSubject) TableName() string {
	return TablePostingSubject
}
