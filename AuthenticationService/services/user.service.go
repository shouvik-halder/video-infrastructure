package services

import (
	"AuthenticationService/dtos"
	"AuthenticationService/helpers"
	"AuthenticationService/interfaces"
	"AuthenticationService/models"
	"errors"
)

type UserServiceImpl struct {
	userRepo interfaces.UserRepository
}

func NewUserService(UR interfaces.UserRepository) interfaces.UserService {
	return &UserServiceImpl{
		userRepo: UR,
	}
}

func (userServ *UserServiceImpl) RegisterService(payload *dtos.UserRegisterDTO) (*models.AuthResponse, error) {

	hashedPassword, err := helpers.HashPassword(payload.Password)
	if err != nil {
		return nil, err
	}
	user, err := userServ.userRepo.Register(payload.Email, hashedPassword)
	if err != nil {
		return nil, err
	}

	token, err := helpers.GenerateToken(user)

	if err != nil {
		return nil, err
	}

	return &models.AuthResponse{
		Token: token,
	}, nil

}

func (userServ *UserServiceImpl) LoginService(payload *dtos.UserLoginDTO) (*models.AuthResponse, error) {
	user, err := userServ.userRepo.Login(payload.Email)
	if err != nil {
		return nil, err
	}

	if ok := helpers.ValidatePassword(payload.Password, user.PasswordHash); !ok {
		return nil, errors.New("Invalid password")
	}

	token, err := helpers.GenerateToken(user)
	if err != nil {
		return nil, err
	}

	return &models.AuthResponse{
		Token: token,
	}, nil
}
