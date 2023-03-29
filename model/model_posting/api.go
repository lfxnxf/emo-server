package model_posting

type AddPostingReq struct {
	Uid        int64    `json:"uid"`
	SubjectIds []int64  `json:"subject_ids"`
	Content    string   `json:"content"`
	Images     []string `json:"images"`
}
