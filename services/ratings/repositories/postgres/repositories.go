package postgres

import (
	"context"
	"database/sql"
	"time"

	"github.com/davq23/jokeapibutbetter/app/data"
)

type Rating struct {
	db *sql.DB
}

func NewRating(db *sql.DB) *Rating {
	return &Rating{
		db: db,
	}
}

func (r *Rating) GetByJokeIDAndUserID(ctx context.Context, jokeID, userID string) (*data.Rating, error) {
	rating := &data.Rating{}
	statement, err := r.db.PrepareContext(
		ctx,
		"SELECT uuid, stars, comment, user_id, joke_id FROM ratings WHERE joke_id = $1 AND user_id = $2 AND deleted_at IS NULL",
	)

	if err != nil {
		return nil, err
	}

	defer statement.Close()

	err = statement.QueryRowContext(ctx, jokeID, userID).Scan(&rating.ID, &rating.Stars, &rating.Comment, &rating.UserID, &rating.JokeID)

	if err != nil {
		return nil, err
	}

	return rating, nil

}

func (r *Rating) GetAllByJokeID(ctx context.Context, jokeID string) ([]*data.Rating, error) {
	var ratings []*data.Rating

	statement, err := r.db.PrepareContext(
		ctx,
		"SELECT uuid, stars, comment, user_id, joke_id FROM ratings WHERE joke_id = $1 AND deleted_at IS NULL",
	)

	if err != nil {
		return nil, err
	}

	defer statement.Close()

	rows, err := statement.QueryContext(ctx, jokeID)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		rating := &data.Rating{}

		err = rows.Scan(&rating.ID, &rating.Stars, &rating.Comment, &rating.UserID, &rating.JokeID)

		if err != nil {
			return nil, err
		}

		ratings = append(ratings, rating)
	}

	return ratings, nil
}

func (r *Rating) GetAllByUserID(ctx context.Context, userID string) ([]*data.Rating, error) {
	var ratings []*data.Rating

	statement, err := r.db.PrepareContext(
		ctx,
		"SELECT uuid, stars, comment, user_id, joke_id FROM ratings WHERE user_id = $1 AND deleted_at IS NULL",
	)

	if err != nil {
		return nil, err
	}

	defer statement.Close()

	rows, err := statement.QueryContext(ctx, userID)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		rating := &data.Rating{}

		err = rows.Scan(&rating.ID, &rating.Stars, &rating.Comment, &rating.UserID, &rating.JokeID)

		if err != nil {
			return nil, err
		}

		ratings = append(ratings, rating)
	}

	return ratings, nil
}

func (r *Rating) GetByID(ctx context.Context, ratingID string) (*data.Rating, error) {
	rating := &data.Rating{}
	statement, err := r.db.PrepareContext(
		ctx,
		"SELECT uuid, stars, comment, user_id, joke_id FROM ratings WHERE uuid = $1 AND deleted_at IS NULL",
	)

	if err != nil {
		return nil, err
	}

	defer statement.Close()

	err = statement.QueryRowContext(ctx, ratingID).Scan(&rating.ID, &rating.Stars, &rating.Comment, &rating.UserID, &rating.JokeID)

	if err != nil {
		return nil, err
	}

	return rating, nil
}

func (r *Rating) Insert(ctx context.Context, rating *data.Rating) error {
	statement, err := r.db.PrepareContext(
		ctx,
		"INSERT INTO ratings (uuid, joke_id, user_id, stars, comment, added_at) VALUES ($1, $2, $3, $4, $5, $6)",
	)

	if err != nil {
		return err
	}

	defer statement.Close()

	now := time.Now()
	rating.AddedAt = &now

	_, err = statement.ExecContext(ctx, rating.ID, rating.JokeID, rating.UserID, rating.Stars, rating.Comment, rating.AddedAt.UTC())

	if err != nil {
		return err
	}

	return nil
}

func (r *Rating) Update(ctx context.Context, rating *data.Rating) error {
	statement, err := r.db.PrepareContext(
		ctx,
		"UPDATE ratings SET user_id = $2, stars = $3, comment = $4, modified_at = $5 WHERE uuid = $1 AND deleted_at IS NULL",
	)

	if err != nil {
		return err
	}

	defer statement.Close()

	now := time.Now()
	rating.ModifiedAt = &now

	_, err = statement.ExecContext(ctx, rating.ID, rating.UserID, rating.Stars, rating.Comment, rating.ModifiedAt.UTC())

	if err != nil {
		return err
	}

	return nil
}
