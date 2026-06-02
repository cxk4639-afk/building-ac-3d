package bootstrap

import (
	"fmt"
	"log"

	"building-ac-3d/backend/internal/config"
	"building-ac-3d/backend/internal/database"
	"building-ac-3d/backend/internal/router"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Application struct {
	cfg    *config.Config
	db     *gorm.DB
	engine *gin.Engine
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
		db:     db,
		engine: engine,
	}, nil
}

func (app *Application) Run() error {
	log.Printf("building-ac-3d backend is running at http://localhost:%s", app.cfg.App.Port)
	return app.engine.Run(app.Addr())
}

func (app *Application) Addr() string {
	return fmt.Sprintf(":%s", app.cfg.App.Port)
}

func (app *Application) DB() *gorm.DB {
	return app.db
}
