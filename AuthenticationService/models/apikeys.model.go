package models

import "time"

type ApiKey struct {
	Id           int64     `json:"id"`
	UserId       int64     `json:"user_id"`
	KeyId        string    `json:"keyId"`
	KeyHash      string    `json:"key_hash"`
	CreatedAt    time.Time `json:"created_at"`
	RevokedAt    time.Time `json:"revoked_at"`
	LastUsedAtAt time.Time `json:"last_used_at"`
}
