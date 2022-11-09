package middlewares

import (
	"context"
	"errors"
	"log"
	"net/http"
	"strings"

	"github.com/davq23/jokeapibutbetter/app/libs"
	"github.com/dgrijalva/jwt-go"
)

type JWTAuth struct {
	logger *log.Logger
	secret string
}

func NewJWTAuth(secret string, logger *log.Logger) *JWTAuth {
	return &JWTAuth{
		secret: secret,
		logger: logger,
	}
}

type AuthClaims struct {
	jwt.StandardClaims
	UserID string
}

type CurrentUserIDKey struct{}

func (j *JWTAuth) AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		formatter := r.Context().Value(FormatterContextKey{}).(libs.Formatter)

		bearerAuthPair := strings.Split(r.Header.Get("Authorization"), " ")

		j.logger.Println(bearerAuthPair)

		if len(bearerAuthPair) != 2 {
			formatter.WriteFormatted(w, libs.StandardReponse{
				Status:  http.StatusForbidden,
				Message: "Forbidden",
			})
			return
		}

		claims := &AuthClaims{}

		token, err := jwt.ParseWithClaims(bearerAuthPair[1], claims, func(token *jwt.Token) (interface{}, error) {
			_, ok := token.Method.(*jwt.SigningMethodHMAC)

			if !ok {
				return nil, errors.New("Forbidden")
			}

			return []byte(j.secret), nil
		})

		if err != nil || !token.Valid {
			if err != nil {
				j.logger.Println(err.Error())
			} else {
				j.logger.Println(token.Claims.Valid().Error())
			}
			formatter.WriteFormatted(w, libs.StandardReponse{
				Status:  http.StatusForbidden,
				Message: "Forbidden",
			})
			return
		}

		next.ServeHTTP(w, r.WithContext(context.WithValue(r.Context(), CurrentUserIDKey{}, claims.UserID)))
	})
}
