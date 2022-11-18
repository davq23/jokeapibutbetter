package main

import (
	"github.com/davq23/jokeapibutbetter/app"
	"github.com/davq23/jokeapibutbetter/services/monolith"
)

func main() {
	appMonolith := &monolith.App{}

	app.RunApp(appMonolith)
}
