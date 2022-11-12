package services

import (
	"context"
	"errors"
	"strings"

	"github.com/davq23/jokeapibutbetter/app/data"
	"github.com/davq23/jokeapibutbetter/app/libs"
	"github.com/davq23/jokeapibutbetter/app/middlewares"
	"github.com/davq23/jokeapibutbetter/services/users/repositories"
	"golang.org/x/crypto/bcrypt"
	"golang.org/x/exp/slices"
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

func (u User) AuthenticateUser(c context.Context, authRequest *libs.AuthRequest) (*data.User, error) {
	user, err := u.FindByEmail(c, authRequest.UsernameOrEmail)

	if err != nil {
		return nil, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Hash), []byte(authRequest.Password))

	if err != nil {
		return nil, err
	}

	return user, nil
}

func (u User) Save(c context.Context, user *data.User) error {
	currentUserID, okCurrent := c.Value(middlewares.CurrentUserIDKey{}).(string)
	currentUser, err := u.userRepository.FetchByID(c, currentUserID)

	if !okCurrent || err != nil {
		return errors.New("currentUserID not found")
	}

	if !slices.Contains(currentUser.Roles, data.USER_ROLE_ADMIN) {
		return errors.New("Forbidden")
	}

	if strings.TrimSpace(user.ID) == "" {
		user.HashPassword()
		user.GenerateID()

		_, err = u.FindByEmail(c, user.Email)

		if err == nil {
			return errors.New("user already in DB")
		}

		return u.userRepository.Insert(c, user)
	}

	return u.userRepository.Update(c, user)
}
