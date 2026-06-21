package models

type ErrorResponse struct {
	Message string
}

type AuthResponse struct {
	Token string `json:"accessToken"`
}

type ApiKeyResponse struct {
	ApiKey string `json:"apiKey"`
}

type ApiKeyVerificationResponse struct {
	UserId int64  `json:"user_id"`
	KeyId  string `json:"key_id"`
}
