package app

import (
    "log"

    "building-ac-3d/backend/internal/config"
    "building-ac-3d/backend/internal/infrastructure/database"
    "building-ac-3d/backend/internal/interface/http/router"
)

type Application struct {
    cfg    *config.Config
    engine *router.Engine
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

    engine := router.New(router.Options{
        Config: cfg,
        DB:     db,
    })

    return &Application{
        cfg:    cfg,
        engine: engine,
    }, nil
}

func (app *Application) Run() error {
    log.Printf("building-ac-3d backend is running at http://localhost:%s", app.cfg.App.Port)
    return app.engine.Run(":" + app.cfg.App.Port)
}
