package app

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/davq23/jokeapibutbetter/app/config"
	"github.com/davq23/jokeapibutbetter/app/data"
	"github.com/davq23/jokeapibutbetter/app/middlewares"
	appServices "github.com/davq23/jokeapibutbetter/app/services"
	"github.com/davq23/jokeapibutbetter/services/jokes/handlers"
	"github.com/davq23/jokeapibutbetter/services/jokes/repositories/postgres"
	"github.com/davq23/jokeapibutbetter/services/jokes/services"
	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
)

type App struct {
	db     *sql.DB
	server *http.Server
}

func (a *App) Shutdown(ctx context.Context) error {
	a.db.Close()
	a.server.Shutdown(ctx)

	return nil
}

func (a *App) Setup() error {
	dbConfig, err := config.LoadDBConfig()
	logger := log.New(os.Stdout, "joke service - ", log.LstdFlags)

	if err != nil {
		return err
	}

	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		dbConfig["host"], dbConfig["port"], dbConfig["user"], dbConfig["password"], dbConfig["dbname"], dbConfig["sslmode"],
	)

	a.db, err = sql.Open("postgres", psqlInfo)

	if err != nil {
		return err
	}

	client := &http.Client{}

	userService := appServices.NewUser("http://host.docker.internal:8080", client, "")
	jokeRepository := postgres.NewJoke(a.db)
	jokeService := services.NewJoke(jokeRepository, userService)
	jokeHandler := handlers.NewJoke(jokeService, logger)
	validate := validator.New()

	bodyValidator := middlewares.NewBodyValidator(validate, func() interface{} {
		return &data.Joke{}
	})

	router := mux.NewRouter()
	router.Use(middlewares.FormatMiddleware)

	postRoutes := router.Methods(http.MethodPost).Subrouter()

	postRoutes.Use(bodyValidator.ValidatorMiddleware)

	getRoutes := router.Methods(http.MethodGet).Subrouter()

	getRoutes.HandleFunc("/hello", jokeHandler.SayHello)

	getRoutes.HandleFunc(
		"/jokes/{id:[a-fA-F0-9]{8}-[a-fA-F0-9]{4}-4[a-fA-F0-9]{3}-[8|9|aA|bB][a-fA-F0-9]{3}-[a-fA-F0-9]{12}}",
		jokeHandler.GetByID,
	)
	postRoutes.HandleFunc("/jokes", jokeHandler.Save)

	a.server = &http.Server{
		Addr:         ":8055",
		Handler:      router,
		WriteTimeout: time.Minute,
		ReadTimeout:  33 * time.Second,
		IdleTimeout:  4 * time.Minute,
	}

	return err
}

func (a *App) Run() error {
	return a.server.ListenAndServe()
}
