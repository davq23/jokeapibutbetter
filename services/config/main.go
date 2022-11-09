package main

import (
	appLib "github.com/davq23/jokeapibutbetter/app"
	"github.com/davq23/jokeapibutbetter/services/config/app"
	_ "github.com/lib/pq"
)

func main() {
	app := &app.App{}

	appLib.RunApp(app)
}
