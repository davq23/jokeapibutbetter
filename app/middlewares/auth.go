package middlewares

import (
	"context"
	"errors"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/davq23/jokeapibutbetter/app/libs"
	"github.com/davq23/jokeapibutbetter/app/utilities"
	"github.com/dgrijalva/jwt-go"
)

type JWTAuth struct {
	logger        *log.Logger
	secret        string
	refreshSecret string
	allowRefresh  bool
}

func NewJWTAuth(secret string, refreshSecret string, allowRefresh bool, logger *log.Logger) *JWTAuth {
	return &JWTAuth{
		secret:        secret,
		refreshSecret: refreshSecret,
		logger:        logger,
		allowRefresh:  allowRefresh,
	}
}

type AuthClaims struct {
	jwt.StandardClaims
	UserID string
}

type CurrentUserIDKey struct{}
type RefreshSecretKey struct{}

func (j JWTAuth) CheckRefreshToken(r *http.Request) (*jwt.Token, error) {
	refreshTokenCookie, err := r.Cookie("refresh")

	if err != nil {
		return nil, err
	}

	if refreshTokenCookie.Expires.After(time.Now()) {
		return nil, errors.New("refresh token cookie is expired")
	}

	claims := &AuthClaims{}

	// Check auth token
	token, err := utilities.ParseJWT(claims, refreshTokenCookie.Value, j.refreshSecret)

	return token, err
}

func (j *JWTAuth) GetCurrentUserMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		bearerAuthPair := strings.Split(r.Header.Get("Authorization"), " ")

		ctx := r.Context()

		if len(bearerAuthPair) == 2 {
			claims := &AuthClaims{}

			// Check auth token
			token, err := utilities.ParseJWT(claims, bearerAuthPair[1], j.secret)

			if err == nil && token.Valid {
				ctx = context.WithValue(ctx, CurrentUserIDKey{}, claims.UserID)
			}
		}

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func (j *JWTAuth) AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		formatter := r.Context().Value(FormatterContextKey{}).(libs.Formatter)

		bearerAuthPair := strings.Split(r.Header.Get("Authorization"), " ")

		if len(bearerAuthPair) != 2 {
			w.WriteHeader(http.StatusForbidden)
			formatter.WriteFormatted(w, libs.StandardReponse{
				Status:  http.StatusForbidden,
				Message: "Forbidden",
			})
			return
		}

		claims := &AuthClaims{}

		// Check auth token
		token, err := utilities.ParseJWT(claims, bearerAuthPair[1], j.secret)

		ctx := r.Context()

		if err != nil || !token.Valid {
			//token, err := j.CheckRefreshToken(r)

			if err != nil || !token.Valid {
				j.logger.Println(err.Error())
				formatter.WriteFormatted(w, libs.StandardReponse{
					Status:  http.StatusForbidden,
					Message: "Forbidden",
				})
				return
			}

			claims = token.Claims.(*AuthClaims)
			ctx = context.WithValue(ctx, RefreshSecretKey{}, true)
		}

		ctx = context.WithValue(ctx, CurrentUserIDKey{}, claims.UserID)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
