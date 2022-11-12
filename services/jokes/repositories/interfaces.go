package repositories

import (
	"context"

	"github.com/davq23/jokeapibutbetter/app/data"
)

type Joke interface {
	GetByID(context.Context, string) (*data.Joke, error)
	GetAll(context.Context, uint64, string, uint64, uint64) ([]*data.Joke, error)
	Insert(context.Context, *data.Joke) error
	Delete(context.Context, *data.Joke) error
	Update(context.Context, *data.Joke) error
}
