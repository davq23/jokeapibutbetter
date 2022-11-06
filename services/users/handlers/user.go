package handlers

import (
	"log"
	"net/http"

	"github.com/davq23/jokeapibutbetter/app/data"
	"github.com/davq23/jokeapibutbetter/app/libs"
	"github.com/davq23/jokeapibutbetter/app/middlewares"
	"github.com/davq23/jokeapibutbetter/services/users/services"
	"github.com/gorilla/mux"
)

type User struct {
	userService *services.User
	logger      *log.Logger
}

func NewUser(userService *services.User, logger *log.Logger) *User {
	return &User{
		userService: userService,
		logger:      logger,
	}
}

func (u *User) SayHello(w http.ResponseWriter, r *http.Request) {
	formatter := r.Context().Value(middlewares.FormatterContextKey{}).(libs.Formatter)

	var aaa map[string]string = make(map[string]string)

	aaa["aaa"] = "akfdkafk"
	aaa["bbb"] = "akfdkafk"

	formatter.WriteFormatted(w, aaa)
}

func (u *User) GetOne(w http.ResponseWriter, r *http.Request) {
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

	user, err := u.userService.FindByID(r.Context(), id)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		u.logger.Println(err.Error())
		formatter.WriteFormatted(w, libs.StandardReponse{
			Status:  http.StatusNotFound,
			Message: "User not found",
		})
		return
	}

	formatter.WriteFormatted(w, libs.StandardReponse{Status: http.StatusOK, Data: user})
}

func (u *User) Save(w http.ResponseWriter, r *http.Request) {
	formatter := r.Context().Value(middlewares.FormatterContextKey{}).(libs.Formatter)

	var user data.User

	err := formatter.ReadFormatted(r.Body, &user)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		u.logger.Println(err.Error())
		formatter.WriteFormatted(w, libs.StandardReponse{
			Status:  http.StatusBadRequest,
			Message: "Invalid request",
		})
		return
	}

	err = u.userService.Save(r.Context(), &user)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		u.logger.Println(err.Error())
		formatter.WriteFormatted(w, libs.StandardReponse{
			Status:  http.StatusBadRequest,
			Message: "Unknown error",
		})
		return
	}

	formatter.WriteFormatted(w, libs.StandardReponse{Status: http.StatusOK, Data: user})
}
