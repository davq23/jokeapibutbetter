package services

import (
	"context"

	"github.com/davq23/jokeapibutbetter/app/data"
	"github.com/davq23/jokeapibutbetter/app/libs"
	"github.com/davq23/jokeapibutbetter/app/services"
	"github.com/davq23/jokeapibutbetter/services/ratings/repositories"
)

type Rating struct {
	jokeService      *services.Joke
	ratingRepository repositories.Rating
	userService      *services.User
}

func NewRating(jokeService *services.Joke, ratingRepository repositories.Rating, userService *services.User) *Rating {
	return &Rating{
		jokeService:      jokeService,
		ratingRepository: ratingRepository,
		userService:      userService,
	}
}

func (rt *Rating) FindByJokeID(ctx context.Context, jokeID string) ([]*data.Rating, error) {
	_, err := rt.jokeService.GetByID(ctx, jokeID, libs.JSONFormatter{})

	if err != nil {
		return nil, err.Err
	}

	return rt.ratingRepository.GetAllByJokeID(ctx, jokeID)
}

func (rt *Rating) FindByUserID(ctx context.Context, userID string) ([]*data.Rating, error) {
	_, err := rt.userService.GetByID(ctx, userID, libs.JSONFormatter{})

	if err != nil {
		return nil, err.Err
	}

	return rt.ratingRepository.GetAllByUserID(ctx, userID)
}

func (rt *Rating) FindByUserIDAndJokeID(ctx context.Context, userID string, jokeID string) (*data.Rating, error) {
	_, err := rt.jokeService.GetByID(ctx, jokeID, libs.JSONFormatter{})

	if err != nil {
		return nil, err.Err
	}

	_, err = rt.userService.GetByID(ctx, userID, libs.JSONFormatter{})

	if err != nil {
		return nil, err.Err
	}

	return rt.ratingRepository.GetByJokeIDAndUserID(ctx, jokeID, userID)
}

func (rt *Rating) Save(ctx context.Context, rating *data.Rating) error {
	_, err := rt.jokeService.GetByID(ctx, rating.JokeID, libs.JSONFormatter{})

	if err != nil {
		return err.Err
	}

	_, err = rt.userService.GetByID(ctx, rating.UserID, libs.JSONFormatter{})

	if err != nil {
		return err.Err
	}

	if rating.ID == "" {
		rating.GenerateID()

		return rt.ratingRepository.Insert(ctx, rating)
	}

	return rt.ratingRepository.Update(ctx, rating)
}
