package main

import (
	appLib "github.com/davq23/jokeapibutbetter/app"
	"github.com/davq23/jokeapibutbetter/services/gateway/app"
)

func main() {
	app := &app.App{}

	appLib.RunApp(app)
}
