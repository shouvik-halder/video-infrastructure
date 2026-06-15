package interfaces

import "AuthenticationService/models"

type ApiKeysRepository interface {
	Create(userId int64, keyId, keyHash string) (*models.ApiKey, error)
	Get(keyId string) (*models.ApiKey, error)
	Revoke(keyId string) (int64, error)
}

type ApiKeysService interface {
	CreateService(userId int64) (*models.ApiKeyResponse, error)
	GetService(keyId string) (*models.ApiKeyResponse, error)
	RevokeService(keyId string) error
}
