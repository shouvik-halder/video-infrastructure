package controllers

import (
	"AuthenticationService/helpers"
	"AuthenticationService/interfaces"
	"AuthenticationService/utils"
	"net/http"
)

type ApiKeysController struct {
	apiKeysService interfaces.ApiKeysService
}

func NewApiKeysController(_apiKeysService interfaces.ApiKeysService) *ApiKeysController {
	return &ApiKeysController{
		apiKeysService: _apiKeysService,
	}
}

func (apiKeysContr *ApiKeysController) CreateController(w http.ResponseWriter, r *http.Request) {
	userId, ok := helpers.GetUserId(r.Context())
	if !ok {
		utils.WriteErrorResponseJson(w, http.StatusUnauthorized, "invalid user")
		return
	}
	response, err := apiKeysContr.apiKeysService.CreateService(userId)
	if err != nil {
		utils.WriteErrorResponseJson(w, http.StatusBadRequest, err.Error())
	}
	utils.WriteResponseJson(w, http.StatusCreated, response)
}

func (apiKeysContr *ApiKeysController) VerifyController(w http.ResponseWriter, r *http.Request) {
	apiKey, ok := helpers.GetApiKey(r.Context())
	if !ok {
		utils.WriteErrorResponseJson(w, http.StatusUnauthorized, "provide valid api key")
		return
	}
	userId, ok := helpers.GetUserId(r.Context())
	if !ok {
		utils.WriteErrorResponseJson(w, http.StatusUnauthorized, "provide valid authorization token")
		return
	}
	response, err := apiKeysContr.apiKeysService.VerifyService(apiKey, userId)
	if err != nil {
		utils.WriteErrorResponseJson(w, http.StatusBadRequest, err.Error())
		return
	}
	utils.WriteResponseJson(w, http.StatusOK, response)
}
