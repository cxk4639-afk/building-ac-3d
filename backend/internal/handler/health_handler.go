package handler

import (
	"gorm.io/gorm"

	"github.com/gin-gonic/gin"
)

type HealthHandler struct {
	db *gorm.DB
}

func NewHealthHandler(db *gorm.DB) *HealthHandler {
	return &HealthHandler{db: db}
}

func (h *HealthHandler) DBHealth(c *gin.Context) {
	sqlDB, err := h.db.DB()
	if err != nil {
		fail(c, 400, err.Error())
		return
	}
	if err := sqlDB.Ping(); err != nil {
		fail(c, 400, err.Error())
		return
	}
	ok(c, map[string]string{"status": "ok"})
}
