package models

type Principal struct {
	UserId int64  `json:"user_id"`
	KeyId  string `json:"key_id"`
}
