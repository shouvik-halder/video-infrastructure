package routers

import (
	"ApiGateway/helpers"
	"ApiGateway/middlewares"
	"net/http"

	"github.com/go-chi/chi/v5"
)

type UploadRouter struct {
	serviceURL string
}

func NewUploadRouter(serviceURL string) *UploadRouter {
	return &UploadRouter{serviceURL: serviceURL}
}

func (ur *UploadRouter) Register(r chi.Router) {
	r.Route("/upload", func(r chi.Router) {

		r.Use(middlewares.AuthenticateApiKey())

		r.Get("/ping", func(w http.ResponseWriter, req *http.Request) {
			helpers.ProxyRequest(w, req, ur.serviceURL, "/ping")
		})
	})
}
