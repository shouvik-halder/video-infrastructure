package routers

import (
	"ApiGateway/helpers"
	"ApiGateway/middlewares"
	"net/http"

	"github.com/go-chi/chi/v5"
)

type AuthRouter struct {
	serviceURL string
	authMiddleware *middlewares.AuthMiddleware
}

func NewAuthRouter(serviceURL string, authMiddleware *middlewares.AuthMiddleware) *AuthRouter {
	return &AuthRouter{
		serviceURL: serviceURL,
		authMiddleware :authMiddleware,
	}
}

func (ar *AuthRouter) Register(r chi.Router) {
	r.Route("/auth", func(r chi.Router) {

		r.Route("/user", func(r chi.Router) {
			r.Post("/register", func(w http.ResponseWriter, req *http.Request) {
				helpers.ProxyRequest(w, req, ar.serviceURL, "/api/v1/user/register")
			})

			r.Post("/login", func(w http.ResponseWriter, req *http.Request) {
				helpers.ProxyRequest(w, req, ar.serviceURL, "/api/v1/user/login")
			})
		})

		r.Group(func(r chi.Router) {
			r.Use(ar.authMiddleware.AuthenticateAccesstoken())

			r.Route("/api-key", func(r chi.Router) {
				r.Post("/create", func(w http.ResponseWriter, req *http.Request) {
					helpers.ProxyRequest(w, req, ar.serviceURL, "/api/v1/api-key/create")
				})

				r.With(ar.authMiddleware.AuthenticateApiKey()).Get("/verify", func(w http.ResponseWriter, req *http.Request) {
					helpers.ProxyRequest(w, req, ar.serviceURL, "/api/v1/api-key/verify")
				})
			})
		})
	})
}