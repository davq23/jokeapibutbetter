package app

import (
	"context"
	"net/http"

	"github.com/davq23/jokeapibutbetter/services/config/handlers"
	"github.com/gorilla/mux"
)

type App struct {
	server *http.Server
}

func (a *App) Shutdown(ctx context.Context) error {
	return nil
}

func (a *App) Run() error {
	return nil
}

func (a *App) Setup() error {
	router := mux.NewRouter()

	configHandler := handlers.NewConfig()

	router.HandleFunc("/config", configHandler.Get).Methods(http.MethodGet)

	a.server = &http.Server{}

	return nil
}
