package services

import (
	"context"
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

func (u Joke) GetByID(c context.Context, id string, formatter libs.Formatter) (*data.Joke, *libs.ApiError) {
	response, err := u.makeRequest(http.MethodGet, "/jokes/"+id, u.authorizationToken, formatter, func() interface{} {
		return &data.Joke{}
	})

	if err != nil {
		return nil, &libs.ApiError{
			Err: err,
		}
	}

	if response.Status != http.StatusOK {
		return nil, &libs.ApiError{
			Err: fmt.Errorf("Error: %d", response.Status),
		}
	}

	joke, ok := response.Data.(*data.Joke)

	if !ok {
		return nil, &libs.ApiError{
			Err: fmt.Errorf("Error: %s", response.Data),
		}
	}

	return joke, nil
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

func (u User) GetByID(c context.Context, id string, formatter libs.Formatter) (*data.User, *libs.ApiError) {
	response, err := u.makeRequest(http.MethodGet, "/users/"+id, u.authorizationToken, formatter, func() interface{} {
		return &data.User{}
	})

	if err != nil {
		return nil, &libs.ApiError{
			Err: err,
		}
	}

	if response.Status != http.StatusOK {
		return nil, &libs.ApiError{
			Err: fmt.Errorf("Error: %d", response.Status),
		}
	}

	user, ok := response.Data.(*data.User)

	if !ok {
		return nil, &libs.ApiError{
			Err: fmt.Errorf("Error: %s", response.Data),
		}
	}

	return user, nil
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
		return nil, err
	}

	return parsedResponse, nil
}
