package routers

import (
	"ApiGateway/helpers"
	"ApiGateway/middlewares"
	"net/http"

	"github.com/go-chi/chi/v5"
)

type AuthRouter struct {
	serviceURL string
}

func NewAuthRouter(serviceURL string) *AuthRouter {
	return &AuthRouter{serviceURL: serviceURL}
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
			r.Use(middlewares.AuthenticateAccesstoken())

			r.Route("/api-key", func(r chi.Router) {
				r.Post("/create", func(w http.ResponseWriter, req *http.Request) {
					helpers.ProxyRequest(w, req, ar.serviceURL, "/api/v1/api-key/create")
				})

				r.With(middlewares.AuthenticateApiKey()).Get("/verify", func(w http.ResponseWriter, req *http.Request) {
					helpers.ProxyRequest(w, req, ar.serviceURL, "/api/v1/api-key/verify")
				})
			})
		})
	})
}