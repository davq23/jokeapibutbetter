package handlers

import (
	"log"
	"net/http"

	"github.com/davq23/jokeapibutbetter/app/data"
	"github.com/davq23/jokeapibutbetter/app/libs"
	"github.com/davq23/jokeapibutbetter/app/middlewares"
	"github.com/davq23/jokeapibutbetter/services/ratings/services"
	"github.com/gorilla/mux"
)

type Rating struct {
	ratingService *services.Rating
	logger        *log.Logger
}

func NewRating(ratingService *services.Rating, logger *log.Logger) *Rating {
	return &Rating{
		ratingService: ratingService,
		logger:        logger,
	}
}

func (rt Rating) GetByUserID(w http.ResponseWriter, r *http.Request) {
	formatter := r.Context().Value(middlewares.FormatterContextKey{}).(libs.Formatter)
	data := mux.Vars(r)
	userID, ok := data["user_id"]

	if !ok {
		w.WriteHeader(http.StatusBadRequest)
		formatter.WriteFormatted(w, libs.StandardReponse{
			Status:  http.StatusBadRequest,
			Message: "Invalid ID",
		})
		return
	}

	ratings, err := rt.ratingService.FindByUserID(r.Context(), userID)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		formatter.WriteFormatted(w, libs.StandardReponse{
			Status:  http.StatusNotFound,
			Message: "An error occurred, probably shitty code",
		})
		return
	}

	formatter.WriteFormatted(w, libs.StandardReponse{Status: http.StatusOK, Data: ratings})
}

func (rt Rating) GetByJokeID(w http.ResponseWriter, r *http.Request) {
	formatter := r.Context().Value(middlewares.FormatterContextKey{}).(libs.Formatter)
	data := mux.Vars(r)
	jokeID, ok := data["joke_id"]

	if !ok {
		w.WriteHeader(http.StatusBadRequest)
		formatter.WriteFormatted(w, libs.StandardReponse{
			Status:  http.StatusBadRequest,
			Message: "Invalid ID",
		})
		return
	}

	ratings, err := rt.ratingService.FindByJokeID(r.Context(), jokeID)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		formatter.WriteFormatted(w, libs.StandardReponse{
			Status:  http.StatusNotFound,
			Message: "An error occurred, probably shitty code",
		})
		return
	}

	formatter.WriteFormatted(w, libs.StandardReponse{Status: http.StatusOK, Data: ratings})
}

func (rt Rating) GetByJokeIDAndUserID(w http.ResponseWriter, r *http.Request) {
	formatter := r.Context().Value(middlewares.FormatterContextKey{}).(libs.Formatter)
	data := mux.Vars(r)
	jokeID, okJokeID := data["joke_id"]
	userID, okUserID := data["user_id"]

	if !okJokeID || !okUserID {
		w.WriteHeader(http.StatusBadRequest)
		formatter.WriteFormatted(w, libs.StandardReponse{
			Status:  http.StatusBadRequest,
			Message: "Invalid ID",
		})
		return
	}

	ratings, err := rt.ratingService.FindByUserIDAndJokeID(r.Context(), userID, jokeID)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		formatter.WriteFormatted(w, libs.StandardReponse{
			Status:  http.StatusNotFound,
			Message: "An error occurred, probably shitty code",
		})
		return
	}

	formatter.WriteFormatted(w, libs.StandardReponse{Status: http.StatusOK, Data: ratings})
}

func (rt Rating) Save(w http.ResponseWriter, r *http.Request) {
	formatter, okFormatter := r.Context().Value(middlewares.FormatterContextKey{}).(libs.Formatter)
	rating, okRating := r.Context().Value(middlewares.ValidatedBodyContextKey{}).(*data.Rating)

	if !okFormatter || !okRating {
		rt.logger.Println("no valid formatter or rating")
		w.WriteHeader(http.StatusBadRequest)
		formatter.WriteFormatted(w, libs.StandardReponse{
			Status:  http.StatusBadRequest,
			Message: "Invalid ID",
		})
		return
	}

	err := rt.ratingService.Save(r.Context(), rating)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		formatter.WriteFormatted(w, libs.StandardReponse{
			Status:  http.StatusNotFound,
			Message: "An error occurred, probably shitty code",
		})
		return
	}

	formatter.WriteFormatted(w, libs.StandardReponse{Status: http.StatusOK, Data: rating})
}
