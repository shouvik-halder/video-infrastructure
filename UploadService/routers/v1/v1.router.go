package v1

import (
	"UploadService/interfaces"

	"github.com/go-chi/chi/v5"
)

type V1Router struct {
	routes []interfaces.Router
}

func NewV1Router(_routes ...interfaces.Router) *V1Router {
	return &V1Router{
		routes: _routes,
	}
}

func (v1 *V1Router) Register(r chi.Router) {
	r.Route("/v1", func(r chi.Router) {
		for _, route := range v1.routes {
			route.Register(r)
		}
	})
}
