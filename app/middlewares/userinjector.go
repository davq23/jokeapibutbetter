package middlewares

import (
	"context"
	"net/http"

	"github.com/davq23/jokeapibutbetter/app/services"
)

type UserInjector struct {
	service services.UserInterface
}

type CurrentUserKey struct{}

func NewUserInjector(us services.UserInterface) *UserInjector {
	return &UserInjector{
		service: us,
	}
}

func (ui UserInjector) UserInjectorMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		currentUserId, ok := r.Context().Value(CurrentUserIDKey{}).(string)
		ctx := r.Context()

		if ok {
			currentUser, err := ui.service.FindByID(r.Context(), currentUserId)

			if err == nil {
				ctx = context.WithValue(r.Context(), CurrentUserKey{}, currentUser)
			} else {
				ctx = context.WithValue(r.Context(), CurrentUserKey{}, nil)
			}
		}

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
