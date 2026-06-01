package main

import (
    "log"

    "building-ac-3d/backend/internal/app"
)

func main() {
    application, err := app.New()
    if err != nil {
        log.Fatalf("init app failed: %v", err)
    }

    if err := application.Run(); err != nil {
        log.Fatalf("run app failed: %v", err)
    }
}
