package services

import (
	"context"
	"database/sql"

	"github.com/davq23/jokeapibutbetter/app/data"
	"github.com/davq23/jokeapibutbetter/app/services"
	"github.com/davq23/jokeapibutbetter/services/ratings/repositories"
)

type Rating struct {
	jokeService      services.JokeInterface
	ratingRepository repositories.Rating
	userService      services.UserInterface
}

func NewRating(jokeService services.JokeInterface, ratingRepository repositories.Rating, userService services.UserInterface) *Rating {
	return &Rating{
		jokeService:      jokeService,
		ratingRepository: ratingRepository,
		userService:      userService,
	}
}

func (rt *Rating) FindByJokeID(ctx context.Context, jokeID string) ([]*data.Rating, error) {
	_, err := rt.jokeService.FindByID(ctx, jokeID)

	if err != nil {
		return nil, err
	}

	return rt.ratingRepository.GetAllByJokeID(ctx, jokeID)
}

func (rt *Rating) FindByUserID(ctx context.Context, userID string) ([]*data.Rating, error) {
	_, err := rt.userService.FindByID(ctx, userID)

	if err != nil {
		return nil, err
	}

	return rt.ratingRepository.GetAllByUserID(ctx, userID)
}

func (rt *Rating) FindByUserIDAndJokeID(ctx context.Context, userID string, jokeID string) (*data.Rating, error) {
	_, err := rt.jokeService.FindByID(ctx, jokeID)

	if err != nil {
		return nil, err
	}

	_, err = rt.userService.FindByID(ctx, userID)

	if err != nil {
		return nil, err
	}

	return rt.ratingRepository.GetByJokeIDAndUserID(ctx, jokeID, userID)
}

func (rt *Rating) Save(ctx context.Context, rating *data.Rating) error {
	currentRating, err := rt.ratingRepository.GetByJokeIDAndUserID(ctx, rating.JokeID, rating.UserID)

	if err != nil {
		if err != sql.ErrNoRows {
			return err
		}

		_, err := rt.jokeService.FindByID(ctx, rating.JokeID)

		if err != nil {
			return err
		}

		_, err = rt.userService.FindByID(ctx, rating.UserID)

		if err != nil {
			return err
		}

		rating.GenerateID()

		return rt.ratingRepository.Insert(ctx, rating)
	}

	rating.ID = currentRating.ID

	return rt.ratingRepository.Update(ctx, rating)
}
