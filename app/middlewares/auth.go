package middlewares

import (
	"context"
	"errors"
	"net/http"
	"strings"

	"github.com/davq23/jokeapibutbetter/app/libs"
	"github.com/dgrijalva/jwt-go"
)

type JWTAuth struct {
	secret string
}

func NewJWTAuth(secret string) *JWTAuth {
	return &JWTAuth{
		secret: secret,
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

		if len(bearerAuthPair) != 2 {
			formatter.WriteFormatted(w, libs.StandardReponse{
				Status:  http.StatusForbidden,
				Message: "Forbidden",
			})
			return
		}

		token, err := jwt.ParseWithClaims(bearerAuthPair[1], &AuthClaims{}, func(token *jwt.Token) (interface{}, error) {
			_, ok := token.Method.(*jwt.SigningMethodECDSA)

			if !ok {
				return nil, errors.New("Forbidden")
			}

			return "", nil
		})

		if err != nil {
			formatter.WriteFormatted(w, libs.StandardReponse{
				Status:  http.StatusForbidden,
				Message: "Forbidden",
			})
			return
		}

		claims, ok := token.Claims.(*AuthClaims)

		if !ok {
			formatter.WriteFormatted(w, libs.StandardReponse{
				Status:  http.StatusForbidden,
				Message: "Forbidden",
			})
			return
		}

		next.ServeHTTP(w, r.WithContext(context.WithValue(r.Context(), CurrentUserIDKey{}, claims.UserID)))
	})
}
