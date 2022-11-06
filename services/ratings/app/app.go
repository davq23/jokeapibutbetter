package app

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/davq23/jokeapibutbetter/app/config"
	"github.com/davq23/jokeapibutbetter/app/data"
	"github.com/davq23/jokeapibutbetter/app/middlewares"
	appServices "github.com/davq23/jokeapibutbetter/app/services"
	"github.com/davq23/jokeapibutbetter/services/ratings/handlers"
	"github.com/davq23/jokeapibutbetter/services/ratings/repositories/postgres"
	"github.com/davq23/jokeapibutbetter/services/ratings/services"
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
	logger := log.New(os.Stdout, "ratings service - ", log.LstdFlags)

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
	jokeService := appServices.NewJoke("http://host.docker.internal:8055", client, "")
	ratingRepository := postgres.NewRating(a.db)
	ratingService := services.NewRating(jokeService, ratingRepository, userService)
	ratingHandler := handlers.NewRating(ratingService, logger)
	validate := validator.New()

	bodyValidator := middlewares.NewBodyValidator(validate, func() interface{} {
		return &data.Rating{}
	})

	router := mux.NewRouter()
	router.Use(middlewares.FormatMiddleware)

	postRoutes := router.Methods(http.MethodPost).Subrouter()
	postRoutes.Use(bodyValidator.ValidatorMiddleware)

	getRoutes := router.Methods(http.MethodGet).Subrouter()

	getRoutes.HandleFunc(
		"/jokes/{joke_id:"+data.IDRegexp+"}/ratings",
		ratingHandler.GetByJokeID,
	)
	getRoutes.HandleFunc(
		"/users/{user_id:"+data.IDRegexp+"}/ratings",
		ratingHandler.GetByUserID,
	)
	getRoutes.HandleFunc(
		"/jokes/{joke_id:"+data.IDRegexp+"}/users/{user_id:"+data.IDRegexp+"}/rating",
		ratingHandler.GetByJokeIDAndUserID,
	)
	postRoutes.HandleFunc("/ratings", ratingHandler.Save)

	a.server = &http.Server{
		Addr:    ":8056",
		Handler: router,
	}

	return err
}

func (a *App) Run() error {
	return a.server.ListenAndServe()
}
