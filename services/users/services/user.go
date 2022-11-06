package services

import (
	"context"
	"errors"
	"strings"

	"github.com/davq23/jokeapibutbetter/app/data"
	"github.com/davq23/jokeapibutbetter/services/users/repositories"
)

type User struct {
	userRepository repositories.User
}

func NewUser(userRepository repositories.User) *User {
	return &User{
		userRepository: userRepository,
	}
}

func (u User) FindByID(c context.Context, id string) (*data.User, error) {
	return u.userRepository.FetchByID(c, id)
}

func (u User) FindByEmail(c context.Context, email string) (*data.User, error) {
	return u.userRepository.FetchByEmail(c, email)
}

func (u User) Save(c context.Context, user *data.User) error {
	if strings.TrimSpace(user.ID) == "" {
		user.HashPassword()
		user.GenerateID()

		_, err := u.FindByEmail(c, user.Email)

		if err == nil {
			return errors.New("User already in DB")
		}

		return u.userRepository.Insert(c, user)
	}

	return u.userRepository.Update(c, user)
}
