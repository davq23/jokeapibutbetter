package app

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

type App interface {
	Setup() error
	Shutdown(ctx context.Context) error
	Run() error
}

func RunApp(app App) {
	err := app.Setup()

	if err != nil {
		log.Fatalln(err.Error())
	}

	go app.Run()

	sigChan := make(chan os.Signal, 5)

	signal.Notify(sigChan, os.Interrupt)
	signal.Notify(sigChan, syscall.SIGTERM)

	sig := <-sigChan

	log.Println("Received terminate, graceful shutdown", sig)

	tc, cancelFunc := context.WithTimeout(context.Background(), 30*time.Second)

	defer cancelFunc()
	defer app.Shutdown(tc)
}
