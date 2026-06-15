package app

import (
	"AuthenticationService/configs"
	"AuthenticationService/configs/db"
	"AuthenticationService/controllers"
	"AuthenticationService/repositories"
	"AuthenticationService/routers"
	v1 "AuthenticationService/routers/v1"
	"AuthenticationService/services"
	"fmt"
	"net/http"
	"time"
)

type Application struct {
	config *configs.Config
	// store *sql.DB
}

func NewApplication() *Application {
	cfg := configs.Load()
	db.SetupDB(cfg)
	return &Application{
		config: cfg,
		// store: ,
	}
}

func (app *Application) Run() error {

	userRepo := repositories.NewUserRepository(db.GetDB())
	userService := services.NewUserService(userRepo)
	userController := controllers.NewUserController(userService)
	userRouter := v1.NewUserRouter(userController)

	apiKeyRepo := repositories.NewApiKeysRepositor(db.GetDB())
	apiKeyService := services.NewApiKeysService(apiKeyRepo)
	apiKeyController := controllers.NewApiKeysController(apiKeyService)
	apiKeyRouter := v1.NewApiKeysRouter(apiKeyController)
	server := &http.Server{
		Addr:         app.config.Server.PORT,
		Handler:      routers.InitializeRouter(userRouter, apiKeyRouter),
		WriteTimeout: 10 * time.Second,
		ReadTimeout:  10 * time.Second,
	}
	fmt.Println("Server started on port", app.config.Server.PORT)
	return server.ListenAndServe()
}
