package services

import (
	"context"
	"errors"

	"github.com/davq23/jokeapibutbetter/app/data"
	"github.com/davq23/jokeapibutbetter/app/middlewares"
	"github.com/davq23/jokeapibutbetter/app/services"
	"github.com/davq23/jokeapibutbetter/services/jokes/repositories"
	"golang.org/x/exp/slices"
)

type Joke struct {
	userService services.UserInterface

	jokeRepository repositories.Joke
}

func NewJoke(jokeRepository repositories.Joke, userService services.UserInterface) *Joke {
	return &Joke{
		jokeRepository: jokeRepository,
		userService:    userService,
	}
}

func (j *Joke) FindAll(c context.Context, limit uint64, language string, userID string, direction uint64, addedAtOffset uint64) ([]*data.Joke, error) {
	if limit == 0 {
		limit = 100
	}

	ctxCopy := c

	currentUser, okCurrentUser := c.Value(middlewares.CurrentUserKey{}).(*data.User)

	if okCurrentUser {
		// Non-users are not allow to see ratings
		if !slices.Contains(currentUser.Roles, data.USER_ROLE_USER) {
			ctxCopy = context.WithValue(ctxCopy, middlewares.CurrentUserIDKey{}, nil)
		}
	}

	return j.jokeRepository.GetAll(ctxCopy, limit, language, userID, direction, addedAtOffset)
}

func (j *Joke) FindByID(c context.Context, id string) (*data.Joke, error) {
	return j.jokeRepository.GetByID(c, id)
}

func (j *Joke) Save(c context.Context, joke *data.Joke) error {
	user, err := j.userService.FindByID(c, joke.AuthorID)

	if err != nil {
		return err
	}

	if joke.ID == "" {
		joke.GenerateID()
		joke.AuthorID = user.ID

		return j.jokeRepository.Insert(c, joke)
	}

	if user.ID != joke.AuthorID {
		return errors.New("invalid user")
	}

	return j.jokeRepository.Update(c, joke)
}
