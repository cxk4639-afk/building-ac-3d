package main

import (
	"log"

	"building-ac-3d/backend/internal/bootstrap"
)

func main() {
	app, err := bootstrap.New()
	if err != nil {
		log.Fatalf("init app failed: %v", err)
	}

	if err := app.Run(); err != nil {
		log.Fatalf("run app failed: %v", err)
	}
}
