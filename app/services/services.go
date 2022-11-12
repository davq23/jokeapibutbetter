package services

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"

	"github.com/davq23/jokeapibutbetter/app/data"
	"github.com/davq23/jokeapibutbetter/app/libs"
)

type Service struct {
	apiUrl             string
	client             *http.Client
	authorizationToken string
}

type GetRequestedData func() interface{}

type Joke struct {
	Service
}

func NewJoke(apiUrl string, client *http.Client, authorizationToken string) *Joke {
	return &Joke{
		Service: Service{
			apiUrl:             apiUrl,
			client:             client,
			authorizationToken: authorizationToken,
		},
	}
}

func (u Joke) FindByID(c context.Context, id string) (*data.Joke, error) {
	formatter := &libs.JSONFormatter{}

	response, err := u.makeRequest(http.MethodGet, "/jokes/"+id, u.authorizationToken, formatter, func() interface{} {
		return &data.Joke{}
	})

	if err != nil {
		return nil, err
	}

	if response.Status != http.StatusOK {
		return nil, fmt.Errorf("error: %d", response.Status)
	}

	joke, ok := response.Data.(*data.Joke)

	if !ok {
		return nil, fmt.Errorf("error: %s", response.Data)
	}

	return joke, nil
}

func (j Joke) Save(c context.Context, joke *data.Joke) error {
	return errors.New("unsupported method")
}

type User struct {
	Service
}

func NewUser(apiUrl string, client *http.Client, authorizationToken string) *User {
	return &User{
		Service: Service{
			apiUrl:             apiUrl,
			client:             client,
			authorizationToken: authorizationToken,
		},
	}
}

func (u User) FindByID(c context.Context, id string) (*data.User, error) {
	formatter := &libs.JSONFormatter{}
	response, err := u.makeRequest(http.MethodGet, "/users/"+id, u.authorizationToken, formatter, func() interface{} {
		return &data.User{}
	})

	if err != nil {
		return nil, err
	}

	if response.Status != http.StatusOK {
		return nil, errors.New("invalid user")
	}

	user, ok := response.Data.(*data.User)

	if !ok {
		return nil, errors.New("invalid user")
	}

	return user, nil
}

func (u User) FindByEmail(c context.Context, email string) (*data.User, error) {
	return nil, errors.New("unsupported method")
}
func (u User) AuthenticateUser(c context.Context, authRequest *libs.AuthRequest) (*data.User, error) {
	return nil, errors.New("unsupported method")
}
func (u User) Save(c context.Context, user *data.User) error {
	return errors.New("unsupported method")
}

func (u Service) makeRequest(method string, path string, token string, formatter libs.Formatter, getRequestedData GetRequestedData) (*libs.StandardReponse, error) {
	requestRawUrl, err := url.JoinPath(u.apiUrl, path)

	if err != nil {
		return nil, err
	}

	requestUrl, err := url.Parse(requestRawUrl)
	query := requestUrl.Query()
	query.Set("format", formatter.GetFormatName())
	requestUrl.RawQuery = query.Encode()

	if err != nil {
		return nil, err
	}

	request := &http.Request{
		Method: method,
		URL:    requestUrl,
		Header: map[string][]string{
			"Authorization": {"Bearer " + token},
		},
	}

	response, err := u.client.Do(request)

	if err != nil {
		return nil, err
	}

	defer response.Body.Close()

	parsedResponse := new(libs.StandardReponse)

	parsedResponse.Data = getRequestedData()

	err = formatter.ReadFormatted(response.Body, parsedResponse)

	if err != nil {
		return parsedResponse, err
	}

	return parsedResponse, nil
}

type Config struct {
	Service
}

func NewConfig(apiUrl string, client *http.Client, authorizationToken string) *Config {
	if apiUrl == "" {
		apiUrl = "http://host.docker.internal:8054"
	}
	return &Config{
		Service: Service{
			apiUrl:             apiUrl,
			client:             client,
			authorizationToken: authorizationToken,
		},
	}
}

func (c *Config) Find(ctx context.Context) (*libs.ConfigResponse, error) {
	parsedResponse, err := c.Service.makeRequest(http.MethodGet, "/config", "", &libs.JSONFormatter{}, func() interface{} {
		return &libs.ConfigResponse{}
	})

	if err != nil {
		return nil, err
	}

	return parsedResponse.Data.(*libs.ConfigResponse), err
}
