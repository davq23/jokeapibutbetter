package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/davq23/jokeapibutbetter/app"
	"github.com/davq23/jokeapibutbetter/app/data"
	"github.com/davq23/jokeapibutbetter/app/libs"
	"github.com/davq23/jokeapibutbetter/app/middlewares"
	jokeHandlers "github.com/davq23/jokeapibutbetter/services/jokes/handlers"
	jokePostgres "github.com/davq23/jokeapibutbetter/services/jokes/repositories/postgres"
	jokeServices "github.com/davq23/jokeapibutbetter/services/jokes/services"
	ratingHandlers "github.com/davq23/jokeapibutbetter/services/ratings/handlers"
	ratingPostgres "github.com/davq23/jokeapibutbetter/services/ratings/repositories/postgres"
	ratingServices "github.com/davq23/jokeapibutbetter/services/ratings/services"
	userHandlers "github.com/davq23/jokeapibutbetter/services/users/handlers"
	userPostgres "github.com/davq23/jokeapibutbetter/services/users/repositories/postgres"
	userServices "github.com/davq23/jokeapibutbetter/services/users/services"
	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

type MonolithicApp struct {
	db     *sql.DB
	server *http.Server
}

func (a *MonolithicApp) Shutdown(ctx context.Context) error {
	a.db.Close()
	a.server.Shutdown(ctx)

	return nil
}

func (a *MonolithicApp) Setup() error {
	godotenv.Load("local.env")

	router := mux.NewRouter()

	frontendRoutes := router.PathPrefix("dashboard").Methods("GET").Subrouter()

	fsHome := http.FileServer(http.Dir("/dist"))

	frontendRoutes.Handle("", fsHome)

	apiRoutes := router.PathPrefix("/api").Subrouter()
	apiRoutes.Use(middlewares.FormatMiddleware)

	config := libs.ConfigResponse{
		JokeServiceURL:    os.Getenv("JOKE_URL"),
		JokeServicePort:   os.Getenv("JOKE_PORT"),
		UserServiceURL:    os.Getenv("USER_URL"),
		UserServicePort:   os.Getenv("USER_PORT"),
		RatingServiceURL:  os.Getenv("RATING_URL"),
		RatingServicePort: os.Getenv("RATING_PORT"),
		DBHost:            os.Getenv("DB_HOST"),
		DBName:            os.Getenv("DB_NAME"),
		DBUser:            os.Getenv("DB_USER"),
		DBPassword:        os.Getenv("DB_PASSWORD"),
		Timezone:          os.Getenv("TIMEZONE"),
		APISecret:         os.Getenv("API_SECRET"),
		SSLMode:           os.Getenv("SSL_MODE"),
	}

	if config.Timezone != "" {
		os.Setenv("TZ", config.Timezone)
	}

	psqlInfo := fmt.Sprintf("postgres://%s:%s@%s/%s?sslmode=%s",
		config.DBUser, config.DBPassword, config.DBHost, config.DBName, config.SSLMode,
	)

	var err error

	a.db, err = sql.Open("postgres", psqlInfo)

	if err != nil {
		return err
	}

	userRepository := userPostgres.NewUser(a.db)
	jokeRepository := jokePostgres.NewJoke(a.db)
	ratingRepository := ratingPostgres.NewRating(a.db)

	userService := userServices.NewUser(userRepository)
	jokeService := jokeServices.NewJoke(jokeRepository, userService)
	ratingService := ratingServices.NewRating(jokeService, ratingRepository, userService)

	authMiddlware := middlewares.NewJWTAuth(config.APISecret, log.New(os.Stdout, "auth middleware -", log.LstdFlags))

	validate := validator.New()

	authBodyValidator := middlewares.NewBodyValidator(validate, func() interface{} {
		return &libs.AuthRequest{}
	})
	userBodyValidator := middlewares.NewBodyValidator(validate, func() interface{} {
		return &data.User{}
	})
	jokeBodyValidator := middlewares.NewBodyValidator(validate, func() interface{} {
		return &data.Joke{}
	})
	ratingBodyValidator := middlewares.NewBodyValidator(validate, func() interface{} {
		return &data.Rating{}
	})

	uh := userHandlers.NewUser(userService, log.New(os.Stdout, "user api -", log.LstdFlags), config.APISecret)
	jh := jokeHandlers.NewJoke(jokeService, log.New(os.Stdout, "joke api -", log.LstdFlags))
	rh := ratingHandlers.NewRating(ratingService, log.New(os.Stdout, "rating api -", log.LstdFlags))

	postRoutes := apiRoutes.Methods(http.MethodPost).Subrouter()
	postRoutes.Use(authMiddlware.AuthMiddleware)

	authPostRoutes := router.Methods("POST").PathPrefix("/auth").Subrouter()
	authPostRoutes.Use(authBodyValidator.ValidatorMiddleware)
	authPostRoutes.HandleFunc("/login", uh.AuthenticateUser)

	userPostRoutes := postRoutes.PathPrefix("/users").Subrouter()
	userPostRoutes.Use(userBodyValidator.ValidatorMiddleware)
	userPostRoutes.HandleFunc("", uh.Save)

	jokePostRoutes := postRoutes.PathPrefix("/jokes").Subrouter()
	jokePostRoutes.Use(jokeBodyValidator.ValidatorMiddleware)
	jokePostRoutes.HandleFunc("", jh.Save)

	ratingPostRoutes := postRoutes.PathPrefix("/ratings").Subrouter()
	ratingPostRoutes.Use(ratingBodyValidator.ValidatorMiddleware)
	ratingPostRoutes.HandleFunc("", rh.Save)

	getRoutes := apiRoutes.Methods(http.MethodGet).Subrouter()

	getRoutes.Handle(
		"/users/whoiam",
		authMiddlware.AuthMiddleware(http.HandlerFunc(uh.CurrentUser)),
	)
	getRoutes.HandleFunc(
		"/jokes/{id:[a-fA-F0-9]{8}-[a-fA-F0-9]{4}-4[a-fA-F0-9]{3}-[8|9|aA|bB][a-fA-F0-9]{3}-[a-fA-F0-9]{12}}",
		jh.GetByID,
	)
	getRoutes.HandleFunc(
		"/jokes",
		jh.GetAll,
	)
	getRoutes.HandleFunc(
		"/users/{id:[a-fA-F0-9]{8}-[a-fA-F0-9]{4}-4[a-fA-F0-9]{3}-[8|9|aA|bB][a-fA-F0-9]{3}-[a-fA-F0-9]{12}}",
		uh.GetOne,
	)
	getRoutes.HandleFunc(
		"/jokes/{joke_id:"+data.IDRegexp+"}/ratings",
		rh.GetByJokeID,
	)
	getRoutes.HandleFunc(
		"/users/{user_id:"+data.IDRegexp+"}/ratings",
		rh.GetByUserID,
	)
	getRoutes.HandleFunc(
		"/jokes/{joke_id:"+data.IDRegexp+"}/users/{user_id:"+data.IDRegexp+"}/rating",
		rh.GetByJokeIDAndUserID,
	)

	a.server = &http.Server{
		Addr:         ":" + os.Getenv("PORT"),
		Handler:      router,
		WriteTimeout: time.Minute,
		ReadTimeout:  2 * time.Second,
		IdleTimeout:  2 * time.Minute,
	}

	return nil
}

func (a *MonolithicApp) Run() error {
	err := a.server.ListenAndServe()

	if err != nil {
		log.Println(err.Error())
	}

	return err
}

func main() {
	a := &MonolithicApp{}

	app.RunApp(a)
}
