package handlers

import (
	"net/http"

	"github.com/davq23/jokeapibutbetter/app/libs"
	"github.com/davq23/jokeapibutbetter/app/middlewares"
)

func Hello(w http.ResponseWriter, r *http.Request) {
	formatter := r.Context().Value(middlewares.FormatterContextKey{}).(libs.Formatter)

	formatter.WriteFormatted(w, libs.StandardReponse{
		Status:  http.StatusOK,
		Message: "API ready",
	})
}
