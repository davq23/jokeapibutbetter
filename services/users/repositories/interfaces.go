package repositories

import (
	"context"

	"github.com/davq23/jokeapibutbetter/app/data"
)

type User interface {
	FetchByID(c context.Context, id string) (*data.User, error)
	FetchByEmail(c context.Context, email string) (*data.User, error)
	Insert(c context.Context, user *data.User) error
	Update(c context.Context, user *data.User) error
}
