package app

import (
	"context"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

type App struct {
	server *http.Server
}

func (a *App) Shutdown(ctx context.Context) error {
	a.server.Shutdown(ctx)
	return nil
}

func (a *App) Setup() error {
	router := mux.NewRouter()

	a.server = &http.Server{
		Handler:      router,
		Addr:         ":8054",
		WriteTimeout: 5 * time.Second,
		ReadTimeout:  time.Second,
	}

	return nil
}

func (a *App) Run() error {
	return a.server.ListenAndServe()
}
