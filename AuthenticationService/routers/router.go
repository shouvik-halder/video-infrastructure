package routers

import (
	"AuthenticationService/interfaces"
	v1 "AuthenticationService/routers/v1"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func InitializeRouter(router ...interfaces.Router) *chi.Mux {
	chiRouter := chi.NewRouter()
	chiRouter.Get("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("pong"))
	})
	v1Router := v1.NewV1Router(router...)
	chiRouter.Route("/api", func(r chi.Router) {
		v1Router.Register(r)
	})
	return chiRouter
}
