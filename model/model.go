package model

const (
	UserTypeHuman = 1 // 自然人
	UserTypeRobot = 2 // 马甲人
)

const (
	OrderDesc = "desc"
	OrderAcs  = "asc"
)

type OrderBy struct {
	Key   string `json:"key"`
	Order string `json:"order"`
}
