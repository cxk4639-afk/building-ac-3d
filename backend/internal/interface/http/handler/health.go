package handler

import (
    "database/sql"
    "net/http"
    "time"

    "building-ac-3d/backend/internal/interface/http/response"
)

type HealthHandler struct {
    db *sql.DB
}

func NewHealthHandler(db *sql.DB) *HealthHandler {
    return &HealthHandler{db: db}
}

func (handler *HealthHandler) DBHealth(w http.ResponseWriter, r *http.Request) {
    if err := handler.db.Ping(); err != nil {
        response.Error(w, http.StatusServiceUnavailable, err.Error())
        return
    }

    response.OK(w, map[string]interface{}{
        "ok":        true,
        "database":  "building_ac_3d",
        "message":   "mysql connected",
        "timestamp": time.Now().Format(time.RFC3339),
    })
}
