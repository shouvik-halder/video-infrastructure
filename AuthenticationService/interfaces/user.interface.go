package interfaces

import (
	"AuthenticationService/dtos"
	"AuthenticationService/models"
)

type UserRepository interface {
	Register(email, password string) (*models.User, error)
	Login(email string) (*models.User, error)
}

type UserService interface {
	RegisterService(payload *dtos.UserRegisterDTO) (*models.AuthResponse, error)
	LoginService(payload *dtos.UserLoginDTO) (*models.AuthResponse, error)
}
