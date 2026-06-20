package routers

import (
	"ApiGateway/gateway"
	"ApiGateway/helpers"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strings"

	"github.com/go-chi/chi/v5"
)

func InitialiseRouters(registry *gateway.ServiceRegistry) *chi.Mux {
	router := chi.NewRouter()
	router.Route("/api", func(r chi.Router) {
		r.HandleFunc("/*", func(w http.ResponseWriter, req *http.Request) {
			prefix := helpers.ExtractPrefix(req.URL.Path) // e.g., "/auth"

			service, ok := registry.Resolve(prefix)
			if !ok {
				http.Error(w, "Service not found", http.StatusNotFound)
				return
			}

			proxyRequest(w, req, service.URL, prefix)
		})
	})
	return router
}

func proxyRequest(w http.ResponseWriter, req *http.Request, targetURL string, prefix string) {
	target, err := url.Parse(targetURL)
	if err != nil {
		http.Error(w, "Invalid service URL", http.StatusInternalServerError)
		return
	}

	proxy := &httputil.ReverseProxy{
		Rewrite: func(proxyReq *httputil.ProxyRequest) {
			proxyReq.SetURL(target)

			originalPath := proxyReq.In.URL.Path
			pathWithoutPrefix := strings.TrimPrefix(originalPath, "/api"+prefix)
			proxyReq.Out.URL.Path = "/api" + pathWithoutPrefix

			proxyReq.SetXForwarded()
			requestID := proxyReq.In.Header.Get("X-Request-ID")
			if requestID == "" {
				requestID = helpers.GenerateRequestID()
			}
			proxyReq.Out.Header.Set("X-Request-ID", requestID)

			proxyReq.Out.Header.Del("X-Internal-Secret")
		},
	}

	proxy.ServeHTTP(w, req)
}
