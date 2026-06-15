package v1

import (
	"AuthenticationService/controllers"
	"AuthenticationService/dtos"
	"AuthenticationService/validator"

	"github.com/go-chi/chi/v5"
)

type UserRouter struct {
	userController *controllers.UserController
}

func NewUserRouter(userControl *controllers.UserController) *UserRouter {
	return &UserRouter{
		userController: userControl,
	}
}

func (userRouter *UserRouter) Register(r chi.Router) {
	r.Route("/user", func(r chi.Router) {
		r.With(validator.ValidateRequest[dtos.UserRegisterDTO]()).Post("/register", userRouter.userController.RegisterController)
		r.With(validator.ValidateRequest[dtos.UserLoginDTO]()).Post("/login", userRouter.userController.LoginController)
	})
}
