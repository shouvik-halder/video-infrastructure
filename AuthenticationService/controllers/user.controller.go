package controllers

import (
	"AuthenticationService/dtos"
	"AuthenticationService/helpers"
	"AuthenticationService/interfaces"
	"AuthenticationService/utils"
	"net/http"
)

type UserController struct {
	userService interfaces.UserService
}

func NewUserController(_userService interfaces.UserService) *UserController {
	return &UserController{
		userService: _userService,
	}
}

func (userContr *UserController) RegisterController(w http.ResponseWriter, r *http.Request) {
	payload, ok := helpers.GetPayload[dtos.UserRegisterDTO](r.Context())
	if !ok {
		utils.WriteErrorResponseJson(w, http.StatusUnprocessableEntity, "invalid json")
		return
	}
	response, err := userContr.userService.RegisterService(payload)
	if err != nil {
		utils.WriteErrorResponseJson(w, http.StatusBadRequest, err.Error())
		return
	}
	utils.WriteResponseJson(w, http.StatusCreated, response)
}

func (userContr *UserController) LoginController(w http.ResponseWriter, r *http.Request) {
	payload, ok:=helpers.GetPayload[dtos.UserLoginDTO](r.Context())
	if !ok{
		utils.WriteErrorResponseJson(w, http.StatusUnprocessableEntity, "invalid payload request")
		return
	}

	response, err:=userContr.userService.LoginService(payload)
	if err!=nil{
		utils.WriteErrorResponseJson(w, http.StatusBadRequest, err.Error())
		return
	}
	utils.WriteResponseJson(w, http.StatusOK, response)
}
