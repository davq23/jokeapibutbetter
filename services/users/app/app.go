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
	"github.com/davq23/jokeapibutbetter/services/users/handlers"
	"github.com/davq23/jokeapibutbetter/services/users/repositories/postgres"
	"github.com/davq23/jokeapibutbetter/services/users/services"
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

	if err != nil {
		return err
	}
	logger := log.New(os.Stdout, "users service - ", log.LstdFlags)

	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		dbConfig["host"], dbConfig["port"], dbConfig["user"], dbConfig["password"], dbConfig["dbname"], dbConfig["sslmode"],
	)

	a.db, err = sql.Open("postgres", psqlInfo)

	if err != nil {
		return err
	}

	userRepository := postgres.NewUser(a.db)
	userService := services.NewUser(userRepository)
	userHandler := handlers.NewUser(userService, logger)
	validate := validator.New()

	bodyValidator := middlewares.NewBodyValidator(validate, func() interface{} {
		return &data.User{}
	})
	jwtMiddleware := middlewares.NewJWTAuth("")

	router := mux.NewRouter()
	router.Use(middlewares.FormatMiddleware)

	postRoutes := router.Methods(http.MethodPost).Subrouter()
	postRoutes.Use(bodyValidator.ValidatorMiddleware)
	postRoutes.Use(jwtMiddleware.AuthMiddleware)

	getRoutes := router.Methods(http.MethodGet).Subrouter()

	getRoutes.HandleFunc(
		"/users/{id:[a-fA-F0-9]{8}-[a-fA-F0-9]{4}-4[a-fA-F0-9]{3}-[8|9|aA|bB][a-fA-F0-9]{3}-[a-fA-F0-9]{12}}",
		userHandler.GetOne,
	)
	getRoutes.HandleFunc("/hello", userHandler.SayHello)
	postRoutes.HandleFunc("/users", userHandler.Save)

	a.server = &http.Server{
		Addr:         ":8080",
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
