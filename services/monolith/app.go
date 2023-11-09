package monolith

import (
	"context"
	"database/sql"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/davq23/jokeapibutbetter/app/data"
	"github.com/davq23/jokeapibutbetter/app/libs"
	"github.com/davq23/jokeapibutbetter/app/middlewares"
	configHandlers "github.com/davq23/jokeapibutbetter/services/config/handlers"
	gatewayHandlers "github.com/davq23/jokeapibutbetter/services/gateway/handlers"
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

	// Asset routes
	router.PathPrefix("/assets").Methods(http.MethodGet).Handler(fsHome)

	// Main routes
	router.PathPrefix("/").Methods(http.MethodGet).Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
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
	}))
}

func (a *App) setupApiRoutes(router *mux.Router, config *libs.ConfigResponse) {
	// Add API prefix
	apiRoutes := router.PathPrefix("/api").Subrouter().StrictSlash(true)

	// Authorization middleware
	authMiddlware := middlewares.NewJWTAuth(config.APISecret, config.RefreshSecret, true, log.New(os.Stdout, "auth middleware -", log.LstdFlags))

	// Body serializer / validation middlwares
	validate := validator.New()
	authBodyValidator := middlewares.NewBodyValidator(validate, func() interface{} { return &libs.AuthRequest{} })
	userBodyValidator := middlewares.NewBodyValidator(validate, func() interface{} { return &data.User{} })
	jokeBodyValidator := middlewares.NewBodyValidator(validate, func() interface{} { return &data.Joke{} })
	ratingBodyValidator := middlewares.NewBodyValidator(validate, func() interface{} { return &data.Rating{} })

	// Global in-out formatting middleware for JSON, XML, YAML
	apiRoutes.Use(middlewares.FormatMiddleware)

	// CORS
	apiRoutes.Use(func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Access-Control-Allow-Headers", "Authorization, Set-Cookie")
			w.Header().Set("Access-Control-Allow-Credentials", "true")
			w.Header().Set("Access-Control-Allow-Methods", strings.Join([]string{http.MethodPost, http.MethodGet}, ","))
			w.Header().Set("Access-Control-Allow-Origin", config.CORSDomains)

			h.ServeHTTP(w, r)
		})
	})

	apiRoutes.Methods(http.MethodOptions).HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	// Repositories
	userRepository := userPostgres.NewUser(a.db)
	jokeRepository := jokePostgres.NewJoke(a.db)
	ratingRepository := ratingPostgres.NewRating(a.db)

	// Services
	userService := userServices.NewUser(userRepository)
	jokeService := jokeServices.NewJoke(jokeRepository, userService)
	ratingService := ratingServices.NewRating(jokeService, ratingRepository, userService)

	userInjector := middlewares.NewUserInjector(userService)
	s3Session := session.Must(session.NewSession(&aws.Config{
		Credentials: credentials.NewStaticCredentials(
			os.Getenv("AWS_ID"),
			os.Getenv("AWS_SECRET"),
			os.Getenv("AWS_TOKEN"),
		),
		Endpoint: aws.String(os.Getenv("AWS_ENDPOINT")),
		Region:   aws.String("us-east-2"),
	}))

	storageManager := libs.NewS3StorageManager(s3.New(s3Session))

	// Handlers
	ch := configHandlers.NewConfig()
	uh := userHandlers.NewUser(storageManager, userService, log.New(os.Stdout, "user api -", log.LstdFlags), config.APISecret, config.RefreshSecret)
	jh := jokeHandlers.NewJoke(jokeService, log.New(os.Stdout, "joke api -", log.LstdFlags))
	rh := ratingHandlers.NewRating(ratingService, log.New(os.Stdout, "rating api -", log.LstdFlags))

	// POST routes
	postRoutes := apiRoutes.Methods(http.MethodPost).Subrouter()
	postRoutes.Use(authMiddlware.AuthMiddleware)

	// - Login (different validator)
	authPostRoutes := apiRoutes.Methods(http.MethodPost).PathPrefix("/auth").Subrouter()
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

	// Get all language codes
	getRoutes.HandleFunc("/config/languages", ch.GetAllLanguages)

	// Get all roles
	getRoutes.HandleFunc("/config/roles", ch.GetAllRoles)

	// Who I Am route
	getRoutes.Handle("/users/whoiam", authMiddlware.AuthMiddleware(http.HandlerFunc(uh.CurrentUser)))

	// User-related routes
	getRoutes.Handle("/users/profile/upload", authMiddlware.AuthMiddleware(http.HandlerFunc(uh.GetUploadProfilePictureURL)))
	getRoutes.HandleFunc("/users/{id:"+data.IDRegexp+"}", uh.GetOne)
	getRoutes.Handle("/users/profile/download", authMiddlware.AuthMiddleware(http.HandlerFunc(uh.GetDownloadProfilePictureURL)))

	// Rating-related routes
	getRoutes.HandleFunc("/users/{user_id:"+data.IDRegexp+"}/ratings", rh.GetByUserID)
	getRoutes.HandleFunc("/jokes/{joke_id:"+data.IDRegexp+"}/users/{user_id:"+data.IDRegexp+"}/rating", rh.GetByJokeIDAndUserID)
	getRoutes.HandleFunc("/jokes/{joke_id:"+data.IDRegexp+"}/ratings", rh.GetByJokeID)

	// Joke-related routes
	getRoutes.HandleFunc("/jokes/{id:"+data.IDRegexp+"}", jh.GetByID)
	getRoutes.Handle("/jokes", authMiddlware.GetCurrentUserMiddleware(
		userInjector.UserInjectorMiddleware(http.HandlerFunc(jh.GetAll)),
	))

	// Hello route
	getRoutes.HandleFunc("/", gatewayHandlers.Hello)
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
		RefreshSecret:     os.Getenv("REFRESH_SECRET"),
		SSLMode:           os.Getenv("SSL_MODE"),
		CORSDomains:       os.Getenv("CORS_DOMAINS"),
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
	//a.setupStaticRoutes(router)

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
	err := a.server.ListenAndServeTLS(os.Getenv("SSL_CERT_FILE"), os.Getenv("SSL_KEY_FILE"))

	if err != nil {
		log.Println(err.Error())
	}

	return err
}
