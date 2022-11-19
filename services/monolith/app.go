package monolith

import (
	"context"
	"database/sql"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"

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

type App struct {
	db     *sql.DB
	server *http.Server
}

func (a *App) Shutdown(ctx context.Context) error {
	a.db.Close()
	a.server.Shutdown(ctx)

	return nil
}

func (a *App) setupStaticRoutes(router *mux.Router) {
	directory := os.DirFS("dist")
	fsHome := http.FileServer(http.FS(directory))

	// Main routes
	router.PathPrefix("/").Methods("GET").Handler(http.StripPrefix("/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		file, err := directory.Open("index.html")

		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			return
		}

		fileStat, err := file.Stat()

		if err != nil {
			w.WriteHeader(http.StatusFailedDependency)
			return
		}

		http.ServeContent(w, r, fileStat.Name(), fileStat.ModTime(), file.(io.ReadSeekCloser))
	})))

	// Asset routes
	router.PathPrefix("/assets").Methods("GET").Handler(fsHome)
}

func (a *App) setupApiRoutes(router *mux.Router, config *libs.ConfigResponse) {
	// Add API prefix
	apiRoutes := router.PathPrefix("/api").Subrouter().StrictSlash(true)

	// Authorization middleware
	authMiddlware := middlewares.NewJWTAuth(config.APISecret, log.New(os.Stdout, "auth middleware -", log.LstdFlags))

	// Body serializer / validation middlwares
	validate := validator.New()
	authBodyValidator := middlewares.NewBodyValidator(validate, func() interface{} { return &libs.AuthRequest{} })
	userBodyValidator := middlewares.NewBodyValidator(validate, func() interface{} { return &data.User{} })
	jokeBodyValidator := middlewares.NewBodyValidator(validate, func() interface{} { return &data.Joke{} })
	ratingBodyValidator := middlewares.NewBodyValidator(validate, func() interface{} { return &data.Rating{} })

	// Global in-out formatting middleware for JSON, XML, YAML
	apiRoutes.Use(middlewares.FormatMiddleware)

	// Repositories
	userRepository := userPostgres.NewUser(a.db)
	jokeRepository := jokePostgres.NewJoke(a.db)
	ratingRepository := ratingPostgres.NewRating(a.db)

	// Services
	userService := userServices.NewUser(userRepository)
	jokeService := jokeServices.NewJoke(jokeRepository, userService)
	ratingService := ratingServices.NewRating(jokeService, ratingRepository, userService)

	// Handlers
	uh := userHandlers.NewUser(userService, log.New(os.Stdout, "user api -", log.LstdFlags), config.APISecret)
	jh := jokeHandlers.NewJoke(jokeService, log.New(os.Stdout, "joke api -", log.LstdFlags))
	rh := ratingHandlers.NewRating(ratingService, log.New(os.Stdout, "rating api -", log.LstdFlags))

	// POST routes
	postRoutes := apiRoutes.Methods(http.MethodPost).Subrouter()
	postRoutes.Use(authMiddlware.AuthMiddleware)

	// - Login (different validator)
	authPostRoutes := router.Methods(http.MethodPost).PathPrefix("/auth").Subrouter()
	authPostRoutes.Use(authBodyValidator.ValidatorMiddleware)
	authPostRoutes.HandleFunc("/login", uh.AuthenticateUser)

	// - Save User
	userPostRoutes := postRoutes.PathPrefix("/users").Subrouter()
	userPostRoutes.Use(userBodyValidator.ValidatorMiddleware)
	userPostRoutes.HandleFunc("", uh.Save)

	// - Save Joke
	jokePostRoutes := postRoutes.PathPrefix("/jokes").Subrouter()
	jokePostRoutes.Use(jokeBodyValidator.ValidatorMiddleware)
	jokePostRoutes.HandleFunc("", jh.Save)

	// - Save Rating
	ratingPostRoutes := postRoutes.PathPrefix("/ratings").Subrouter()
	ratingPostRoutes.Use(ratingBodyValidator.ValidatorMiddleware)
	ratingPostRoutes.HandleFunc("", rh.Save)

	// GET routes
	getRoutes := apiRoutes.Methods(http.MethodGet).Subrouter()

	// Who I Am route
	getRoutes.Handle("/users/whoiam", authMiddlware.AuthMiddleware(http.HandlerFunc(uh.CurrentUser)))

	// User-related routes
	getRoutes.HandleFunc("/users/{id:"+data.IDRegexp+"}", uh.GetOne)

	// Rating-related routes
	getRoutes.HandleFunc("/users/{user_id:"+data.IDRegexp+"}/ratings", rh.GetByUserID)
	getRoutes.HandleFunc("/jokes/{joke_id:"+data.IDRegexp+"}/users/{user_id:"+data.IDRegexp+"}/rating", rh.GetByJokeIDAndUserID)
	getRoutes.HandleFunc("/jokes/{joke_id:"+data.IDRegexp+"}/ratings", rh.GetByJokeID)

	// Joke-related routes
	getRoutes.HandleFunc("/jokes/{id:"+data.IDRegexp+"}", jh.GetByID)
	getRoutes.HandleFunc("/jokes", jh.GetAll)
}

func (a *App) Setup() error {
	godotenv.Load("local.env")

	// Router
	router := mux.NewRouter().StrictSlash(true)

	// Configuration
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

	// API routes
	a.setupApiRoutes(router, &config)

	// Static routes
	a.setupStaticRoutes(router)

	a.server = &http.Server{
		Addr:         ":" + os.Getenv("PORT"),
		Handler:      router,
		WriteTimeout: 30 * time.Second,
		ReadTimeout:  2 * time.Second,
		IdleTimeout:  2 * time.Minute,
	}

	return nil
}

func (a *App) Run() error {
	err := a.server.ListenAndServe()

	if err != nil {
		log.Println(err.Error())
	}

	return err
}
