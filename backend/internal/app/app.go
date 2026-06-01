package app

import (
    "fmt"
    "log"
    "net/http"

    "building-ac-3d/backend/internal/config"
    "building-ac-3d/backend/internal/infrastructure/database"
    "building-ac-3d/backend/internal/interface/http/router"
)

type Application struct {
    cfg    *config.Config
    server *http.Server
}

func New() (*Application, error) {
    cfg := config.Load()

    db, err := database.Open(cfg.Database)
    if err != nil {
        return nil, err
    }

    if err := database.Migrate(db); err != nil {
        return nil, err
    }

    handler := router.New(router.Options{
        Config: cfg,
        DB:     db,
    })

    return &Application{
        cfg: cfg,
        server: &http.Server{
            Addr:    ":" + cfg.App.Port,
            Handler: handler,
        },
    }, nil
}

func (app *Application) Run() error {
    log.Printf("building-ac-3d backend is running at http://localhost:%s", app.cfg.App.Port)
    return app.server.ListenAndServe()
}

func (app *Application) Addr() string {
    return fmt.Sprintf(":%s", app.cfg.App.Port)
}
