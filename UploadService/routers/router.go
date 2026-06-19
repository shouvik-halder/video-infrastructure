package routers

import (
	"UploadService/interfaces"
	v1 "UploadService/routers/v1"

	"github.com/go-chi/chi/v5"
)

func InitializeRouters(routers ...interfaces.Router) *chi.Mux {
	chiRouter := chi.NewRouter()

	v1Router := v1.NewV1Router(routers...)
	chiRouter.Route("/api", func(r chi.Router) {
		v1Router.Register(r)
	})
	return chiRouter
}
