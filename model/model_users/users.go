package model_users

import "time"

const (
	TableUsers = "users"
)

const (
	GenderMale   = 1 // 男
	GenderFemale = 2 // 女
)

func GetGenderText(gender int64) string {
	return map[int64]string{
		GenderMale:   "男",
		GenderFemale: "女",
	}[gender]
}

// Users  用户表
type Users struct {
	ID           int64     `gorm:"column:id" json:"id"`                                 //  自增id
	Nickname     string    `gorm:"column:nickname" json:"nickname"`                     //  昵称
	Gender       int64     `gorm:"column:gender" json:"gender"`                         //  性别，1：男，2：女
	Phone        string    `gorm:"column:phone" json:"phone"`                           //  手机号
	CountryCode  string    `gorm:"column:country_code" json:"country_code"`             //  手机号前缀 eg:86
	Password     string    `gorm:"column:password" json:"password"`                     //  密码
	Birthday     string    `gorm:"column:birthday" json:"birthday"`                     //  生日
	Portrait     string    `gorm:"column:portrait" json:"portrait"`                     //  头像
	Introduction string    `gorm:"column:introduction" json:"introduction"`             //  简介
	Token        string    `gorm:"column:token" json:"token"`                           //  token
	UserType     int64     `gorm:"column:user_type" json:"user_type"`                   //  1：自然用户，2：马甲用户
	LoginTime    int64     `gorm:"column:login_time" json:"login_time"`                 //  登录时间
	CreateTime   time.Time `gorm:"column:create_time; default:null" json:"create_time"` //  创建时间
}

func (Users) TableName() string {
	return TableUsers
}
