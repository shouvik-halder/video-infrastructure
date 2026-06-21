package routers

import (
	"ApiGateway/helpers"
	"ApiGateway/middlewares"
	"net/http"

	"github.com/go-chi/chi/v5"
)

type UploadRouter struct {
	serviceURL     string
	authMiddleware *middlewares.AuthMiddleware
}

func NewUploadRouter(serviceURL string, authMiddleware *middlewares.AuthMiddleware) *UploadRouter {
	return &UploadRouter{
		serviceURL:     serviceURL,
		authMiddleware: authMiddleware,
	}
}

func (ur *UploadRouter) Register(r chi.Router) {
	r.Route("/upload", func(r chi.Router) {
		r.Use(ur.authMiddleware.AuthenticateAccesstoken())
		r.Use(ur.authMiddleware.AuthenticateApiKey())

		r.Get("/ping", func(w http.ResponseWriter, req *http.Request) {
			helpers.ProxyRequest(w, req, ur.serviceURL, "/ping")
		})
	})
}
