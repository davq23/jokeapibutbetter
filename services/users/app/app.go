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
	"github.com/davq23/jokeapibutbetter/app/libs"
	"github.com/davq23/jokeapibutbetter/app/middlewares"
	appServices "github.com/davq23/jokeapibutbetter/app/services"
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
	logger := log.New(os.Stdout, "users service - ", log.LstdFlags)
	configService := appServices.NewConfig("http://localhost:8054", &http.Client{}, os.Getenv("INTERNAL_TOKEN"))
	config, err := configService.Find(context.Background())
	if err != nil {
		return err
	}
	config.FixValues()
	os.Setenv("TZ", config.Timezone)
	logger.Println("Configuration completed")

	psqlInfo := fmt.Sprintf("postgres://%s:%s@%s/%s?sslmode=%s",
		config.DBUser, config.DBPassword, config.DBHost, config.DBName, config.SSLMode,
	)

	a.db, err = sql.Open("postgres", psqlInfo)

	if err != nil {
		return err
	}

	userRepository := postgres.NewUser(a.db)
	userService := services.NewUser(userRepository)
	userHandler := handlers.NewUser(userService, logger, config.APISecret, config.RefreshSecret)
	validate := validator.New()

	userBodyValidator := middlewares.NewBodyValidator(validate, func() interface{} {
		return &data.User{}
	})
	authBodyValidator := middlewares.NewBodyValidator(validate, func() interface{} {
		return &libs.AuthRequest{}
	})

	jwtMiddleware := middlewares.NewJWTAuth(config.APISecret, config.RefreshSecret, true, logger)

	router := mux.NewRouter()
	router.Use(middlewares.FormatMiddleware)

	authRoutes := router.PathPrefix("/auth").Methods(http.MethodPost).Subrouter()
	authRoutes.Use(authBodyValidator.ValidatorMiddleware)
	authRoutes.HandleFunc("/users/authenticate", userHandler.AuthenticateUser)

	postRoutes := router.Methods(http.MethodPost).Subrouter()
	postRoutes.Use(jwtMiddleware.AuthMiddleware)
	postRoutes.Use(userBodyValidator.ValidatorMiddleware)

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
	err := a.server.ListenAndServe()

	log.Println(err)

	return err
}
