package models

import "github.com/golang-jwt/jwt/v5"

type Claims struct {
	UserID int64  `json:"user_id"`
	Email  string `json:"email"`
	jwt.RegisteredClaims
}

type validateApiKeyResponse struct {
	UserId int64  `json:"user_id"`
	KeyId  string `json:"key_id"`
}
