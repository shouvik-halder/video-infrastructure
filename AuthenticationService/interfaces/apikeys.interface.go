package interfaces

import "AuthenticationService/models"

type ApiKeysRepository interface {
	Create(userId int64, keyId, keyHash string) (*models.ApiKey, error)
	Get(keyId string, userId int64) (*models.ApiKey, error)
	Revoke(keyId string) (int64, error)
}

type ApiKeysService interface {
	CreateService(userId int64) (*models.ApiKeyResponse, error)
	VerifyService(keyId string, userId int64) (*models.ApiKeyVerificationResponse, error)
	RevokeService(keyId string) error
}
