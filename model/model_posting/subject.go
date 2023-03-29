package model_posting

import "time"

const (
	TableSubject = "subject"

	SubjectStatusNormal  = 1
	SubjectStatusDeleted = 101
)

// Subject 话题表
type Subject struct {
	ID         int64     `gorm:"column:id" json:"id"`                                 //自增id
	Name       string    `gorm:"column:name" json:"name"`                             //话题名称
	Status     int64     `gorm:"column:status" json:"status"`                         //状态，1:正常，101:删除
	CreateTime time.Time `gorm:"column:create_time; default:null" json:"create_time"` //创建时间
}

func (s *Subject) TableName() string {
	return TableSubject
}
