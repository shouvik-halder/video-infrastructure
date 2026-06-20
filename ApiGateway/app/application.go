package app

import (
	"ApiGateway/configs"
	"ApiGateway/gateway"
	"ApiGateway/routers"
	"net/http"
	"time"
)

type Application struct {
	config          *configs.Config
	serviceRegistry *gateway.ServiceRegistry
	routeRegistry   *gateway.RouteRegistry
}

func NewApplication() *Application {
	cfg := configs.Load()

	_serviceRegistry := gateway.NewServiceRegistry()
	_serviceRegistry.RegisterAll(cfg.Service.AUTH_SERVICE_URL, cfg.Service.UPLOAD_SERVICE_URL)

	_routeRegistry := gateway.NewRouteRegistry()
	_routeRegistry.RegisterAll()

	return &Application{
		config:          cfg,
		serviceRegistry: _serviceRegistry,
		routeRegistry:   _routeRegistry,
	}
}

func (app *Application) Run() error {

	server := &http.Server{
		Addr:         app.config.Server.PORT,
		Handler:      routers.InitialiseRouters(app.serviceRegistry, app.routeRegistry),
		WriteTimeout: 10 * time.Second,
		ReadTimeout:  10 * time.Second,
	}

	server.ListenAndServe()
	return nil
}
