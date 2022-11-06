package middlewares

import (
	"context"
	"net/http"

	"github.com/davq23/jokeapibutbetter/app/libs"
	"github.com/go-playground/validator/v10"
)

type DataGetter func() interface{}

type ValidatedBodyContextKey struct{}

type BodyValidator struct {
	validator *validator.Validate
	GetData   DataGetter
}

func NewBodyValidator(validator *validator.Validate, getData DataGetter) *BodyValidator {
	return &BodyValidator{
		validator: validator,
		GetData:   getData,
	}
}

func (bv BodyValidator) ValidatorMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		formatter := r.Context().Value(FormatterContextKey{}).(libs.Formatter)

		body := bv.GetData()

		err := formatter.ReadFormatted(r.Body, body)

		if err != nil {
			formatter.WriteFormatted(w, libs.StandardReponse{
				Status:  http.StatusUnprocessableEntity,
				Message: err.Error(),
			})
			return
		}

		err = bv.validator.Struct(body)

		if err != nil {
			formatter.WriteFormatted(w, libs.StandardReponse{
				Status:  http.StatusBadRequest,
				Message: err.Error(),
			})
			return
		}

		next.ServeHTTP(w, r.WithContext(context.WithValue(r.Context(), ValidatedBodyContextKey{}, body)))
	})

}
