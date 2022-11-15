package postgres

import (
	"context"
	"database/sql"
	"time"

	"github.com/davq23/jokeapibutbetter/app/data"
)

type Joke struct {
	db *sql.DB
}

func NewJoke(db *sql.DB) *Joke {
	return &Joke{
		db: db,
	}
}

func (j *Joke) GetByID(c context.Context, id string) (*data.Joke, error) {
	joke := &data.Joke{}
	statement, err := j.db.PrepareContext(
		c,
		"SELECT j.uuid, j.text, j.author_id, j.description, j.lang, j.added_at, u.username, u.email FROM jokes j JOIN users u ON j.author_id = u.uuid WHERE j.uuid = $1 AND j.deleted_at IS NULL",
	)

	if err != nil {
		return nil, err
	}

	defer statement.Close()

	joke.User = &data.User{}

	err = statement.QueryRowContext(c, id).Scan(&joke.ID, &joke.Text, &joke.AuthorID, &joke.Description, &joke.Language, &joke.AddedAt, &joke.User.Username, &joke.User.Email)

	joke.User.ID = joke.AuthorID

	if err != nil {
		return nil, err
	}

	return joke, nil
}

func (j *Joke) GetAll(c context.Context, limit uint64, language string, direction uint64, addedAtOffset uint64) ([]*data.Joke, error) {
	sqlSentence := "SELECT j.uuid, j.text, j.author_id, j.description, j.lang, j.added_at, u.username, u.email FROM jokes j JOIN users u ON j.author_id = u.uuid WHERE"
	sqlSentence += " j.deleted_at IS NULL"

	if len(language) > 2 {
		if language != "" {
			sqlSentence += " AND j.lang = $1"
		} else {
			sqlSentence += " AND j.lang != $1"
		}
	} else {
		if language != "" {
			sqlSentence += " AND j.lang LIKE CONCAT(cast($1 as varchar(6)), '%')"
		} else {
			sqlSentence += " AND j.lang != $1"
		}
	}

	if direction == 0 {
		sqlSentence += " AND j.added_at >= TO_TIMESTAMP($2)"
	} else {
		sqlSentence += " AND j.added_at < TO_TIMESTAMP($2)"
	}
	sqlSentence += " ORDER BY j.added_at DESC"
	statement, err := j.db.PrepareContext(c, sqlSentence)

	if err != nil {
		return nil, err
	}

	defer statement.Close()

	rows, err := statement.QueryContext(c, language, addedAtOffset)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var jokes []*data.Joke = make([]*data.Joke, 0, limit)

	for rows.Next() {
		joke := &data.Joke{}

		joke.User = &data.User{}

		err = rows.Scan(&joke.ID, &joke.Text, &joke.AuthorID, &joke.Description, &joke.Language, &joke.AddedAt, &joke.User.Username, &joke.User.Email)

		if err != nil {
			return nil, err
		}

		jokes = append(jokes, joke)
	}

	return jokes, nil
}

func (j *Joke) Insert(c context.Context, joke *data.Joke) error {
	statement, err := j.db.PrepareContext(
		c,
		"INSERT INTO jokes (uuid, text, author_id, description, lang, added_at) VALUES ($1, $3, $2, $4, $5, $6)",
	)

	if err != nil {
		return err
	}

	defer statement.Close()

	now := time.Now()
	joke.AddedAt = &now

	_, err = statement.ExecContext(c, joke.ID, joke.AuthorID, joke.Text, joke.Description, joke.Language, joke.AddedAt.UTC())

	if err != nil {
		return err
	}

	return nil
}

func (j *Joke) Delete(c context.Context, joke *data.Joke) error {
	return nil
}

func (j *Joke) Update(c context.Context, joke *data.Joke) error {
	statement, err := j.db.PrepareContext(
		c,
		"UPDATE jokes SET text = $2, author_id = $1, description = $3, lang = $4, modified_at = $5 WHERE uuid = $6 AND deleted_at IS NULL",
	)

	if err != nil {
		return err
	}

	defer statement.Close()

	now := time.Now()
	joke.ModifiedAt = &now

	_, err = statement.ExecContext(c, joke.AuthorID, joke.Text, joke.Description, joke.Language, joke.ModifiedAt.UTC(), joke.ID)

	if err != nil {
		return err
	}

	return nil
}
