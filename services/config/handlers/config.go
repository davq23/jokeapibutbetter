package handlers

import (
	"net/http"
	"os"

	"github.com/davq23/jokeapibutbetter/app/libs"
	"github.com/davq23/jokeapibutbetter/app/middlewares"
)

type Config struct {
}

func NewConfig() *Config {
	return &Config{}
}

func (c *Config) Get(w http.ResponseWriter, r *http.Request) {
	formatter, ok := r.Context().Value(middlewares.FormatterContextKey{}).(libs.Formatter)

	if !ok {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

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

	formatter.WriteFormatted(w, libs.StandardReponse{
		Status: http.StatusOK,
		Data:   config,
	})
}
