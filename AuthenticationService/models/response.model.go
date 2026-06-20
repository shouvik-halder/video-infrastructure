package models

type ErrorResponse struct {
	Message string
}

type AuthResponse struct {
	Token string `json:"accessToken"`
}

type ApiKeyResponse struct {
	ApiKeyId string `json:"apiKeyId"`
	UserId   int64  `json:"userId"`
}
