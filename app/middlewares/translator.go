package middlewares

import (
	"net/http"

	"golang.org/x/text/language"
)

func TranslatorMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		values := r.URL.Query()

		if values.Has("lang") {
			switch values.Get("lang") {
			case language.English.String():
			}
		} else {
			return
		}
	})
}
