package app

import (
	"ApiGateway/configs"
	"ApiGateway/gateway"
	"ApiGateway/models"
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

	registry.Register("api/auth",
		models.Service{
			Name: "auth-service",
			URL:  cfg.Service.AUTH_SERVICE_URL,
		})

	registry.Register("api/upload",
		models.Service{
			Name: "upload-service",
			URL:  cfg.Service.UPLOAD_SERVICE_URL,
		})
	return &Application{
		config:          cfg,
		serviceRegistry: registry,
	}
}

func (app *Application) Run() error {

	server := &http.Server{
		Addr:         app.config.Server.PORT,
		Handler:      nil,
		WriteTimeout: 10 * time.Second,
		ReadTimeout:  10 * time.Second,
	}

	server.ListenAndServe()
	return nil
}
