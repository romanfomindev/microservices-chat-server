package main

import (
	"context"
	"github.com/romanfomindev/microservices-chat-server/internal/app"
	"log"
)

//var configPath string = ".env"

func main() {
	ctx := context.Background()

	a, err := app.NewApp(ctx)
	if err != nil {
		log.Fatalf("failed to init app: %s", err.Error())
	}

	err = a.Run()
	if err != nil {
		log.Fatalf("failed to run app: %s", err.Error())
	}
}
