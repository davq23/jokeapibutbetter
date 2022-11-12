package app

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

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
	client := &http.Client{}
	logger := log.New(os.Stdout, "ratings service - ", log.LstdFlags)
	configService := appServices.NewConfig("http://localhost:8054", client, os.Getenv("INTERNAL_TOKEN"))
	config, err := configService.Find(context.Background())
	os.Setenv("TZ", config.Timezone)
	if err != nil {
		return err
	}
	config.FixValues()
	logger.Println("Configuration completed")
	psqlInfo := fmt.Sprintf("postgres://%s:%s@%s/%s?sslmode=%s",
		config.DBUser, config.DBPassword, config.DBHost, config.DBName, config.SSLMode,
	)

	a.db, err = sql.Open("postgres", psqlInfo)

	if err != nil {
		return err
	}

	userService := appServices.NewUser(config.UserServiceURL+":"+config.UserServicePort, client, "")
	jokeService := appServices.NewJoke(config.JokeServiceURL+":"+config.JokeServicePort, client, "")
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
		Addr:         ":8056",
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
