package app

import (
	"UploadService/configs"
	"UploadService/routers"
	"fmt"
	"net/http"
	"time"
)

type Application struct {
	config *configs.Config
}

func NewApplication() *Application {
	cfg := configs.Load()
	return &Application{
		config: cfg,
	}
}

func (app *Application) Run() error {

	server := &http.Server{
		Addr:         app.config.Server.PORT,
		Handler:      routers.InitializeRouters(),
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	fmt.Println("Starting server on Port ", app.config.Server.PORT)

	return server.ListenAndServe()
}
