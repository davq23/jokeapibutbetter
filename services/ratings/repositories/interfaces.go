package repositories

import (
	"context"

	"github.com/davq23/jokeapibutbetter/app/data"
)

type Rating interface {
	GetAllByJokeID(ctx context.Context, jokeID string) ([]*data.Rating, error)
	GetAllByUserID(ctx context.Context, userID string) ([]*data.Rating, error)
	GetByJokeIDAndUserID(ctx context.Context, jokeID, userID string) (*data.Rating, error)
	GetByID(ctx context.Context, ratingID string) (*data.Rating, error)
	Insert(ctx context.Context, rating *data.Rating) error
	Update(ctx context.Context, rating *data.Rating) error
}
