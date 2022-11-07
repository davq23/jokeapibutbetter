package main

import (
	appLib "github.com/davq23/jokeapibutbetter/app"
	"github.com/davq23/jokeapibutbetter/services/config/app"
)

func main() {
	app := &app.App{}

	appLib.RunApp(app)
}
