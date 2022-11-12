package handlers

import (
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/davq23/jokeapibutbetter/app/data"
	"github.com/davq23/jokeapibutbetter/app/libs"
	"github.com/davq23/jokeapibutbetter/app/middlewares"
	"github.com/davq23/jokeapibutbetter/app/services"
	"github.com/gorilla/mux"
)

type Joke struct {
	jokeService services.JokeInterface
	logger      *log.Logger
}

func NewJoke(jokeService services.JokeInterface, logger *log.Logger) *Joke {
	return &Joke{
		jokeService: jokeService,
		logger:      logger,
	}
}

func (j *Joke) SayHello(w http.ResponseWriter, r *http.Request) {
	client := &http.Client{}

	response, err := client.Get("http://host.docker.internal:8080/hello")

	if err != nil {
		io.WriteString(w, err.Error())
		return
	}

	defer response.Body.Close()

	bytes, err := ioutil.ReadAll(response.Body)

	if err != nil {
		io.WriteString(w, err.Error())
		return
	}

	w.Write(bytes)
}

func (j *Joke) GetByID(w http.ResponseWriter, r *http.Request) {
	formatter := r.Context().Value(middlewares.FormatterContextKey{}).(libs.Formatter)
	data := mux.Vars(r)
	id, ok := data["id"]

	if !ok {
		w.WriteHeader(http.StatusBadRequest)
		formatter.WriteFormatted(w, libs.StandardReponse{
			Status:  http.StatusBadRequest,
			Message: "Invalid ID",
		})
		return
	}

	user, err := j.jokeService.FindByID(r.Context(), id)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		j.logger.Println(err.Error())
		formatter.WriteFormatted(w, libs.StandardReponse{
			Status:  http.StatusBadRequest,
			Message: "An error occured",
		})
		return
	}

	formatter.WriteFormatted(w, libs.StandardReponse{Status: http.StatusOK, Data: user})
}

func (j *Joke) GetAll(w http.ResponseWriter, r *http.Request) {
	formatter := r.Context().Value(middlewares.FormatterContextKey{}).(libs.Formatter)
	limit, _ := strconv.ParseUint(r.URL.Query().Get("limit"), 10, 64)
	addedAtOffsetString := r.URL.Query().Get("offset")
	addedAtOffset, errOffset := strconv.ParseUint(addedAtOffsetString, 10, 64)
	direction, _ := strconv.ParseUint(r.URL.Query().Get("direction"), 10, 64)
	language := r.URL.Query().Get("lang")

	var err error

	if addedAtOffsetString == "" {
		addedAtOffset = 0
	} else if errOffset != nil {
		addedAtTimeOffset, err := time.Parse(time.RFC3339, addedAtOffsetString)

		if err != nil {
			formatter.WriteFormatted(w, libs.StandardReponse{
				Status:  http.StatusBadRequest,
				Message: "Invalid offset",
			})
			return
		}

		addedAtOffset = uint64(addedAtTimeOffset.Unix())
	}

	jokes, err := j.jokeService.FindAll(r.Context(), limit, language, direction, addedAtOffset)

	if err != nil {
		formatter.WriteFormatted(w, libs.StandardReponse{
			Status:  http.StatusBadRequest,
			Message: err.Error(),
		})
		return
	}

	formatter.WriteFormatted(w, libs.StandardReponse{
		Status: http.StatusOK,
		Data:   jokes,
	})
}

func (j *Joke) Save(w http.ResponseWriter, r *http.Request) {
	formatter, okFormatter := r.Context().Value(middlewares.FormatterContextKey{}).(libs.Formatter)

	joke, okJoke := r.Context().Value(middlewares.ValidatedBodyContextKey{}).(*data.Joke)

	if !okFormatter || !okJoke {
		formatter.WriteFormatted(w, libs.StandardReponse{
			Status:  http.StatusBadRequest,
			Message: "Invalid request",
		})
		return
	}

	err := j.jokeService.Save(r.Context(), joke)

	if err != nil {
		formatter.WriteFormatted(w, libs.StandardReponse{
			Status:  http.StatusBadRequest,
			Message: err.Error(),
		})
		return
	}

	formatter.WriteFormatted(w, libs.StandardReponse{
		Status: http.StatusOK,
		Data:   joke,
	})
}
