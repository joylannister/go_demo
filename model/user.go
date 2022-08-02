package model

type UserInfo struct {
	Id       string `json:"id" db:"id"`
	UserName string `json:"user_name" db:"user_name"`
	NickName string `json:"nickname" db:"nickname"`
}
