package services

import (
	"context"
	"errors"

	"github.com/davq23/jokeapibutbetter/app/data"
	"github.com/davq23/jokeapibutbetter/app/services"
	"github.com/davq23/jokeapibutbetter/services/jokes/repositories"
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
