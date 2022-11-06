package handlers

import (
	"io"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/davq23/jokeapibutbetter/app/data"
	"github.com/davq23/jokeapibutbetter/app/libs"
	"github.com/davq23/jokeapibutbetter/app/middlewares"
	"github.com/davq23/jokeapibutbetter/services/jokes/services"
	"github.com/gorilla/mux"
)

type Joke struct {
	jokeService *services.Joke
	logger      *log.Logger
}

func NewJoke(jokeService *services.Joke, logger *log.Logger) *Joke {
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
