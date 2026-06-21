package services

import (
	"AuthenticationService/helpers"
	"AuthenticationService/interfaces"
	"AuthenticationService/models"
	"fmt"
	"strings"
)

type ApiKeysServiceImpl struct {
	apiRepo interfaces.ApiKeysRepository
}

func NewApiKeysService(ar interfaces.ApiKeysRepository) interfaces.ApiKeysService {
	return &ApiKeysServiceImpl{
		apiRepo: ar,
	}
}

func (apiServ *ApiKeysServiceImpl) CreateService(userId int64) (*models.ApiKeyResponse, error) {

	keyID, err := helpers.GenerateAlphaNum(16)
	if err != nil {
		return nil, err
	}

	keySecret, err := helpers.GenerateAlphaNum(64)
	if err != nil {
		return nil, err
	}

	apikeyStr := fmt.Sprintf("ak_%s_%s", keyID, keySecret)
	keyHash := helpers.HashKey(keySecret)

	_, dberr := apiServ.apiRepo.Create(userId, keyID, keyHash)
	if dberr != nil {
		return nil, dberr
	}

	return &models.ApiKeyResponse{
		ApiKey: apikeyStr,
	}, nil
}

func (apiServ *ApiKeysServiceImpl) VerifyService(key string, userId int64) (*models.ApiKeyVerificationResponse, error) {
	parts := strings.Split(key, "_")
	apiKey, err := apiServ.apiRepo.Get(parts[1], userId)
	if err != nil {
		return nil, err
	}

	if ok := helpers.VerifyKey(parts[2], apiKey.KeyHash); !ok {
		return nil, fmt.Errorf("invalid api key")
	}
	return &models.ApiKeyVerificationResponse{
		UserId: userId,
		KeyId:  parts[1],
	}, nil
}

func (apiServ *ApiKeysServiceImpl) RevokeService(keyId string) error {
	// rowsAffected, err := apiServ.apiRepo.Revoke(keyId)
	// if err != nil {
	// 	return err
	// }

	// if rowsAffected == 0 {
	// 	return fmt.Errorf("api key not found or already revoked")
	// }

	return nil
}
