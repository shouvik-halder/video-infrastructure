package routers

import (
	"ApiGateway/gateway"
	"ApiGateway/helpers"
	"ApiGateway/middlewares"
	"net/http"
	"net/http/httputil"
	"net/url"

	"github.com/go-chi/chi/v5"
)

func InitialiseRouters(
	serviceRegistry *gateway.ServiceRegistry,
	routeRegistry *gateway.RouteRegistry,
) *chi.Mux {

	router := chi.NewRouter()

	router.Route("/api", func(r chi.Router) {

		r.HandleFunc("/*", func(w http.ResponseWriter, req *http.Request) {

			route, ok := routeRegistry.Resolve(
				req.Method,
				req.URL.Path,
			)

			if !ok {
				http.Error(w, "route not found", http.StatusNotFound)
				return
			}

			// Apply authentication middleware if route requires it
			if route.RequireAuth {
				authMiddleware := middlewares.AuthenticateApiKey()
				authMiddleware(http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
					proxyToService(w, req, serviceRegistry, route)
				})).ServeHTTP(w, req)
				return
			}

			proxyToService(w, req, serviceRegistry, route)
		})
	})

	return router
}

// proxyToService resolves the service and proxies the request
func proxyToService(
	w http.ResponseWriter,
	req *http.Request,
	serviceRegistry *gateway.ServiceRegistry,
	route *gateway.RouteMapping,
) {
	service, ok := serviceRegistry.Resolve(route.ServiceKey)
	if !ok {
		http.Error(w, "service not found", http.StatusBadGateway)
		return
	}

	proxyRequest(
		w,
		req,
		service.URL,
		route.TargetPath,
	)
}

func proxyRequest(
	w http.ResponseWriter,
	req *http.Request,
	targetURL string,
	targetPath string,
) {

	target, err := url.Parse(targetURL)
	if err != nil {
		http.Error(w, "Invalid service URL", http.StatusInternalServerError)
		return
	}

	proxy := &httputil.ReverseProxy{
		Rewrite: func(proxyReq *httputil.ProxyRequest) {

			proxyReq.SetURL(target)

			proxyReq.Out.URL.Path = targetPath

			proxyReq.SetXForwarded()

			requestID := proxyReq.In.Header.Get("X-Request-ID")
			if requestID == "" {
				requestID = helpers.GenerateRequestID()
			}

			proxyReq.Out.Header.Set(
				"X-Request-ID",
				requestID,
			)

			proxyReq.Out.Header.Del(
				"X-Internal-Secret",
			)
		},
	}

	proxy.ServeHTTP(w, req)
}
