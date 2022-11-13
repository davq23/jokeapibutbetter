package middlewares

import (
	"context"
	"net/http"
	"net/textproto"

	"github.com/davq23/jokeapibutbetter/app/libs"
)

type FormatterContextKey struct{}

func FormatMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		values := r.URL.Query()

		var formatter libs.Formatter

		if values.Has("format") {
			switch values.Get("format") {
			case libs.JSON_FORMAT:
				w.Header().Set(textproto.CanonicalMIMEHeaderKey("Content-Type"), "application/json")
				formatter = &libs.JSONFormatter{}
			case libs.XML_FORMAT:
				w.Header().Set(textproto.CanonicalMIMEHeaderKey("Content-Type"), "application/xml")
				formatter = &libs.XMLFormatter{}
			case libs.YAML_FORMAT:
				w.Header().Set(textproto.CanonicalMIMEHeaderKey("Content-Type"), "application/x-yaml")
				formatter = &libs.YAMLFormatter{}
			default:
				w.Header().Set(textproto.CanonicalMIMEHeaderKey("Content-Type"), "text/plain")
				w.WriteHeader(http.StatusBadRequest)
				w.Write([]byte("Unsupported format"))
				return
			}
		} else {
			w.Header().Set(textproto.CanonicalMIMEHeaderKey("Content-Type"), "application/json")
			formatter = &libs.JSONFormatter{}
		}

		next.ServeHTTP(
			w,
			r.WithContext(context.WithValue(r.Context(), FormatterContextKey{}, formatter)),
		)
	})
}

func ReadFormattedRequestBodyMiddleware(next http.Handler, data interface{}) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		formatter, ok := r.Context().Value(FormatterContextKey{}).(libs.Formatter)

		if !ok {
			w.Header().Add("Content-Type", "text/plain")
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("Unsupported format"))
			return
		}

		mapper := make(map[string]string)

		if formatter.ReadFormatted(r.Body, mapper) != nil {
			w.Header().Add("Content-Type", "text/plain")
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("Unsupported format"))
			return
		}

		next.ServeHTTP(
			w,
			r,
		)
	})
}
