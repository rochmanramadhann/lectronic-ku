package models

type Token struct {
	ID     uint   `json:"id"`
	UserID uint   `json:"user_id"`
	Token  string `json:"token"`
}
