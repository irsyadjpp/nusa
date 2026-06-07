package main

import (
	"log"

	"github.com/nusa/backend/internal/bootstrap"
)

func main() {
	app, err := bootstrap.New()
	if err != nil {
		log.Fatalf("Failed to initialize application: %v", err)
	}

	if err := app.Run(); err != nil {
		log.Fatalf("Failed to run application: %v", err)
	}
}
