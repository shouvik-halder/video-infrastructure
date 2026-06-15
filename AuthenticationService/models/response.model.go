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
