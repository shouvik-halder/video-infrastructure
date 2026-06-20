package routers

import (
	"UploadService/interfaces"
	v1 "UploadService/routers/v1"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func InitializeRouters(routers ...interfaces.Router) *chi.Mux {
	chiRouter := chi.NewRouter()

	chiRouter.Get("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("pong"))
	})

	v1Router := v1.NewV1Router(routers...)
	chiRouter.Route("/api", func(r chi.Router) {
		v1Router.Register(r)
	})
	return chiRouter
}
