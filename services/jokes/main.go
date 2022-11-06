package main

import (
	appLib "github.com/davq23/jokeapibutbetter/app"
	"github.com/davq23/jokeapibutbetter/services/jokes/app"
	_ "github.com/lib/pq"
)

func main() {
	jokeApp := &app.App{}

	appLib.RunApp(jokeApp)
}
