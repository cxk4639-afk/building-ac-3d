package router

import (
    "database/sql"
    "net/http"

    "building-ac-3d/backend/internal/config"
    "building-ac-3d/backend/internal/interface/http/handler"
    "building-ac-3d/backend/internal/interface/http/middleware"
)

type Options struct {
    Config *config.Config
    DB     *sql.DB
}

func New(options Options) http.Handler {
    mux := http.NewServeMux()

    authHandler := handler.NewAuthHandler(options.DB)
    systemHandler := handler.NewSystemHandler(options.DB)
    healthHandler := handler.NewHealthHandler(options.DB)

    mux.HandleFunc("/api/db/health", healthHandler.DBHealth)
    mux.HandleFunc("/api/auth/login", authHandler.Login)
    mux.HandleFunc("/api/auth/me", authHandler.Me)
    mux.HandleFunc("/api/auth/logout", authHandler.Logout)

    mux.HandleFunc("/api/system/tenants", systemHandler.Tenants)
    mux.HandleFunc("/api/system/users", systemHandler.Users)
    mux.HandleFunc("/api/system/roles", systemHandler.Roles)
    mux.HandleFunc("/api/system/menus", systemHandler.Menus)

    return middleware.CORS(options.Config.CORS.AllowOrigin)(mux)
}
