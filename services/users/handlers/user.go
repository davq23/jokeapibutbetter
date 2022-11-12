package handlers

import (
	"log"
	"net/http"
	"time"

	"github.com/davq23/jokeapibutbetter/app/data"
	"github.com/davq23/jokeapibutbetter/app/libs"
	"github.com/davq23/jokeapibutbetter/app/middlewares"
	"github.com/davq23/jokeapibutbetter/app/services"
	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/mux"
)

type User struct {
	jwtSecret   string
	userService services.UserInterface
	logger      *log.Logger
}

func NewUser(userService services.UserInterface, logger *log.Logger, jwtSecret string) *User {
	return &User{
		jwtSecret:   jwtSecret,
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

func (u *User) AuthenticateUser(w http.ResponseWriter, r *http.Request) {
	formatter, okFormatter := r.Context().Value(middlewares.FormatterContextKey{}).(libs.Formatter)
	auth, okAuth := r.Context().Value(middlewares.ValidatedBodyContextKey{}).(*libs.AuthRequest)

	if !okAuth || !okFormatter {
		w.WriteHeader(http.StatusBadRequest)
		u.logger.Println("Valid user or formatter body not found")
		formatter.WriteFormatted(w, libs.StandardReponse{
			Status:  http.StatusBadRequest,
			Message: "Invalid request",
		})
		return
	}

	user, err := u.userService.AuthenticateUser(r.Context(), auth)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		u.logger.Println(err.Error())
		formatter.WriteFormatted(w, libs.StandardReponse{
			Status:  http.StatusBadRequest,
			Message: "Invalid request",
		})
		return
	}

	claims := middlewares.AuthClaims{
		UserID: user.ID,
		StandardClaims: jwt.StandardClaims{
			IssuedAt:  time.Now().Unix(),
			ExpiresAt: time.Now().Add(10 * time.Minute).Unix(),
		},
	}

	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenClaims.SignedString([]byte(u.jwtSecret))

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		u.logger.Println(err.Error())
		formatter.WriteFormatted(w, libs.StandardReponse{
			Status:  http.StatusInternalServerError,
			Message: "Invalid request",
		})
		return
	}

	formatter.WriteFormatted(w, libs.StandardReponse{Status: http.StatusOK, Data: libs.AuthResponse{
		Token:    token,
		UserID:   user.ID,
		Username: user.Username,
	}})
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
	formatter, okFormatter := r.Context().Value(middlewares.FormatterContextKey{}).(libs.Formatter)
	user, okUser := r.Context().Value(middlewares.ValidatedBodyContextKey{}).(*data.User)

	if !okUser || !okFormatter {
		w.WriteHeader(http.StatusBadRequest)
		u.logger.Println("Valid user or formatter body not found")
		formatter.WriteFormatted(w, libs.StandardReponse{
			Status:  http.StatusBadRequest,
			Message: "Invalid request",
		})
		return
	}

	err := u.userService.Save(r.Context(), user)

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
