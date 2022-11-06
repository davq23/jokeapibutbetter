package main

import (
	appLib "github.com/davq23/jokeapibutbetter/app"
	"github.com/davq23/jokeapibutbetter/services/ratings/app"
	_ "github.com/lib/pq"
)

func main() {
	ratingApp := &app.App{}

	appLib.RunApp(ratingApp)
}
