package app

import (
	"context"
	"net/http"
	"time"

	"github.com/davq23/jokeapibutbetter/app/middlewares"
	"github.com/davq23/jokeapibutbetter/services/config/handlers"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

type App struct {
	server *http.Server
}

func (a *App) Shutdown(ctx context.Context) error {
	a.server.Shutdown(ctx)

	return nil
}

func (a *App) Run() error {
	return a.server.ListenAndServe()
}

func (a *App) Setup() error {
	godotenv.Load("local.env")

	router := mux.NewRouter()

	configHandler := handlers.NewConfig()

	router.Use(middlewares.FormatMiddleware)

	router.HandleFunc("/config", configHandler.Get).Methods(http.MethodGet)

	a.server = &http.Server{
		Addr:         ":8054",
		Handler:      router,
		WriteTimeout: time.Minute,
		ReadTimeout:  time.Second,
		IdleTimeout:  time.Minute,
	}

	return nil
}
