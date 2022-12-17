package handlers

import (
	"net/http"

	"github.com/davq23/jokeapibutbetter/app/data"
	"github.com/davq23/jokeapibutbetter/app/libs"
	"github.com/davq23/jokeapibutbetter/app/middlewares"
	"golang.org/x/text/language"
)

type Config struct {
}

func NewConfig() *Config {
	return &Config{}
}

func (c *Config) GetAllLanguages(w http.ResponseWriter, r *http.Request) {
	formatter, _ := r.Context().Value(middlewares.FormatterContextKey{}).(libs.Formatter)
	formatter.WriteFormatted(w, libs.StandardReponse{
		Status: http.StatusOK,
		Data: []string{
			language.English.String(),
			language.French.String(),
			language.EuropeanSpanish.String(),
		},
	})
}

func (c *Config) GetAllRoles(w http.ResponseWriter, r *http.Request) {
	formatter, _ := r.Context().Value(middlewares.FormatterContextKey{}).(libs.Formatter)
	formatter.WriteFormatted(w, libs.StandardReponse{
		Status: http.StatusOK,
		Data: []string{
			data.USER_ROLE_ADMIN,
			data.USER_ROLE_USER,
			data.USER_ROLE_RSS,
		},
	})
}
