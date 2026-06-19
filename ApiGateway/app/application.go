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
}

func NewApplication() *Application {
	cfg := configs.Load()

	registry := gateway.NewServiceRegistry()
	registry.RegisterAll(cfg.Service.AUTH_SERVICE_URL, cfg.Service.UPLOAD_SERVICE_URL)

	return &Application{
		config:          cfg,
		serviceRegistry: registry,
	}
}

func (app *Application) Run() error {

	server := &http.Server{
		Addr:         app.config.Server.PORT,
		Handler:      routers.InitialiseRouters(app.serviceRegistry),
		WriteTimeout: 10 * time.Second,
		ReadTimeout:  10 * time.Second,
	}

	server.ListenAndServe()
	return nil
}
