package postgres

import (
	"context"
	"database/sql"
	"time"

	"github.com/davq23/jokeapibutbetter/app/data"
)

type User struct {
	db *sql.DB
}

func NewUser(db *sql.DB) *User {
	return &User{
		db: db,
	}
}

func (u *User) Insert(c context.Context, user *data.User) error {
	statement, err := u.db.PrepareContext(
		c,
		"INSERT INTO users (uuid, username, email, hash, added_at) VALUES ($1, $2, $3, $4, $5)",
	)

	if err != nil {
		return err
	}

	defer statement.Close()

	now := time.Now()
	user.AddedAt = &now

	_, err = statement.ExecContext(c, user.ID, user.Username, user.Email, user.Hash, user.AddedAt.UTC())

	if err != nil {
		return err
	}

	return nil
}

func (u *User) Update(c context.Context, user *data.User) error {
	statement, err := u.db.PrepareContext(
		c,
		"UPDATE users SET username = $1, email = $2, hash = $3 , modified_at = $4 WHERE uuid = $5",
	)

	if err != nil {
		return err
	}

	defer statement.Close()

	now := time.Now()
	user.ModifiedAt = &now

	_, err = statement.ExecContext(c, user.Username, user.Email, user.Hash, user.ModifiedAt.UTC(), user.ID)

	if err != nil {
		return err
	}

	return nil
}

func (u *User) FetchByEmail(c context.Context, email string) (*data.User, error) {
	user := &data.User{}
	statement, err := u.db.PrepareContext(
		c,
		"SELECT uuid, username, email, hash FROM users WHERE email = $1 AND deleted_at IS NULL",
	)

	if err != nil {
		return nil, err
	}

	defer statement.Close()

	err = statement.QueryRowContext(c, email).Scan(&user.ID, &user.Username, &user.Email, &user.Hash)

	if err != nil {
		return nil, err
	}

	return user, nil
}

func (u *User) FetchByID(c context.Context, id string) (*data.User, error) {
	user := &data.User{}
	statement, err := u.db.PrepareContext(
		c,
		"SELECT uuid, username, email, hash FROM users WHERE uuid = $1 AND deleted_at IS NULL",
	)

	if err != nil {
		return nil, err
	}

	defer statement.Close()

	err = statement.QueryRowContext(c, id).Scan(&user.ID, &user.Username, &user.Email, &user.Hash)

	if err != nil {
		return nil, err
	}

	return user, nil
}
