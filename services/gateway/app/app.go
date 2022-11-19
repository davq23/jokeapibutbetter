package app

import (
	"context"
	"net/http"
	"os"
	"time"

	appServices "github.com/davq23/jokeapibutbetter/app/services"
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
	configService := appServices.NewConfig(os.Getenv("CONFIG_URL"), &http.Client{}, os.Getenv("INTERNAL_TOKEN"))
	config, err := configService.Find(context.Background())
	if err != nil {
		return err
	}

	os.Setenv("TZ", config.Timezone)
	//logger := log.New(os.Stdout, "gateway service - ", log.LstdFlags)

	router := mux.NewRouter()

	a.server = &http.Server{
		Handler:      router,
		Addr:         ":8080",
		WriteTimeout: 5 * time.Second,
		ReadTimeout:  time.Second,
	}

	return nil
}

func (a *App) Run() error {
	return a.server.ListenAndServe()
}
