package services

import (
	"context"

	"github.com/davq23/jokeapibutbetter/app/data"
	"github.com/davq23/jokeapibutbetter/app/libs"
)

type UserInterface interface {
	FindByID(c context.Context, id string) (*data.User, error)
	FindByEmail(c context.Context, email string) (*data.User, error)
	AuthenticateUser(c context.Context, authRequest *libs.AuthRequest) (*data.User, error)
	Save(c context.Context, user *data.User) error
}

type JokeInterface interface {
	FindByID(c context.Context, id string) (*data.Joke, error)
	Save(c context.Context, joke *data.Joke) error
}

type RatingsInterface interface {
	FindByJokeID(ctx context.Context, jokeID string) ([]*data.Rating, error)
	FindByUserID(ctx context.Context, userID string) ([]*data.Rating, error)
	FindByUserIDAndJokeID(ctx context.Context, userID string, jokeID string) (*data.Rating, error)
	Save(ctx context.Context, rating *data.Rating) error
}
