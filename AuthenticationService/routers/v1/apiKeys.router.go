package v1

import (
	"AuthenticationService/controllers"
	"AuthenticationService/middlewares"

	"github.com/go-chi/chi/v5"
)

type ApiKeysRouter struct{
	apiKeysController *controllers.ApiKeysController
}

func NewApiKeysRouter(apiKeysControl *controllers.ApiKeysController) *ApiKeysRouter {
	return &ApiKeysRouter{
		apiKeysController: apiKeysControl,
	}
}

func (apiKeysRouter *ApiKeysRouter)Register(r chi.Router){
	r.Route("/api-key", func(r chi.Router) {
		r.With(middlewares.ValidateToken()).Post("/create", apiKeysRouter.apiKeysController.CreateController)
	})
} 