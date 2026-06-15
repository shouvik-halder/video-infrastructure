package services

import (
	"AuthenticationService/helpers"
	"AuthenticationService/interfaces"
	"AuthenticationService/models"
	"fmt"
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
	keyHash:= helpers.HashKey(keySecret)

	_ , dberr:=apiServ.apiRepo.Create(userId,keyID, keyHash)
	if dberr!=nil{
		return nil, err
	}

	return &models.ApiKeyResponse{
		ApiKey: apikeyStr,
	}, nil
}

func (apiServ *ApiKeysServiceImpl) GetService(keyId string) (*models.ApiKeyResponse, error) {
	return nil, nil
}

func (apiServ *ApiKeysServiceImpl) RevokeService(keyId string) error {
	return nil
}
