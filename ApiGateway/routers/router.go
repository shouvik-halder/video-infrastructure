package routers

import (
	"ApiGateway/interfaces"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func InitialiseRouters(
	routers ...interfaces.Router,
) *chi.Mux {

	chiRouter := chi.NewRouter()
	chiRouter.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("content-type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"status"}:"ok"`))
	})

	chiRouter.Route("/api", func(r chi.Router) {
		for _, route := range routers {
			route.Register(r)
		}
	})

	return chiRouter
}
