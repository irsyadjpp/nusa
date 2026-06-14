package main

import (
	"github.com/nusa/backend/internal/bootstrap"
)

func main() {
	app, err := bootstrap.New()
	if err != nil {
		panic(err)
	}

	if err := app.Run(); err != nil {
		panic(err)
	}
}
