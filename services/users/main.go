package main

import (
	appLib "github.com/davq23/jokeapibutbetter/app"
	"github.com/davq23/jokeapibutbetter/services/users/app"
	_ "github.com/lib/pq"
)

func main() {
	userApp := &app.App{}

	appLib.RunApp(userApp)
}
