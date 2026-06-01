package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

type App struct {
	db *sql.DB
}

type HealthResponse struct {
	OK        bool   `json:"ok"`
	Database  string `json:"database"`
	Message   string `json:"message"`
	Timestamp string `json:"timestamp"`
}

type VisualConfig struct {
	ID          int64           `json:"id"`
	Name        string          `json:"name"`
	Description string          `json:"description"`
	ConfigData  json.RawMessage `json:"configData"`
	CreatedAt   string          `json:"createdAt"`
	UpdatedAt   string          `json:"updatedAt"`
}

type CreateVisualConfigRequest struct {
	Name        string          `json:"name"`
	Description string          `json:"description"`
	ConfigData  json.RawMessage `json:"configData"`
}

func main() {
	loadEnvFile()

	db, err := openDB()
	if err != nil {
		log.Fatalf("connect mysql failed: %v", err)
	}
	defer db.Close()

	app := &App{db: db}
	if err := app.migrate(); err != nil {
		log.Fatalf("migrate database failed: %v", err)
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/api/db/health", app.withCORS(app.handleDBHealth))
	mux.HandleFunc("/api/visual-configs", app.withCORS(app.handleVisualConfigs))

	port := env("APP_PORT", "8080")
	log.Printf("building-ac-3d backend is running at http://localhost:%s", port)
	log.Fatal(http.ListenAndServe(":"+port, mux))
}

func loadEnvFile() {
	if err := godotenv.Load(); err == nil {
		return
	}

	_ = godotenv.Load("backend/.env")
}

func openDB() (*sql.DB, error) {
	user := env("MYSQL_USER", "root")
	password := env("MYSQL_PASSWORD", "")
	host := env("MYSQL_HOST", "127.0.0.1")
	port := env("MYSQL_PORT", "3306")
	database := env("MYSQL_DATABASE", "building_ac_3d")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=true&loc=Local", user, password, host, port, database)
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}

	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(5)
	db.SetConnMaxLifetime(time.Hour)

	if err := db.Ping(); err != nil {
		_ = db.Close()
		return nil, err
	}

	return db, nil
}

func (app *App) migrate() error {
	_, err := app.db.Exec(`
CREATE TABLE IF NOT EXISTS visual_configs (
  id BIGINT PRIMARY KEY AUTO_INCREMENT,
  name VARCHAR(100) NOT NULL,
  description VARCHAR(255) NOT NULL DEFAULT '',
  config_data JSON NOT NULL,
  created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
`)
	return err
}

func (app *App) handleDBHealth(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		writeError(w, http.StatusMethodNotAllowed, "method not allowed")
		return
	}

	if err := app.db.Ping(); err != nil {
		writeJSON(w, http.StatusServiceUnavailable, HealthResponse{
			OK:        false,
			Database:  env("MYSQL_DATABASE", "building_ac_3d"),
			Message:   err.Error(),
			Timestamp: time.Now().Format(time.RFC3339),
		})
		return
	}

	writeJSON(w, http.StatusOK, HealthResponse{
		OK:        true,
		Database:  env("MYSQL_DATABASE", "building_ac_3d"),
		Message:   "mysql connected",
		Timestamp: time.Now().Format(time.RFC3339),
	})
}

func (app *App) handleVisualConfigs(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		app.listVisualConfigs(w, r)
	case http.MethodPost:
		app.createVisualConfig(w, r)
	default:
		writeError(w, http.StatusMethodNotAllowed, "method not allowed")
	}
}

func (app *App) listVisualConfigs(w http.ResponseWriter, r *http.Request) {
	rows, err := app.db.Query(`
SELECT id, name, description, config_data, created_at, updated_at
FROM visual_configs
ORDER BY updated_at DESC, id DESC
`)
	if err != nil {
		writeError(w, http.StatusInternalServerError, err.Error())
		return
	}
	defer rows.Close()

	configs := make([]VisualConfig, 0)
	for rows.Next() {
		var item VisualConfig
		var configData []byte
		var createdAt time.Time
		var updatedAt time.Time

		if err := rows.Scan(&item.ID, &item.Name, &item.Description, &configData, &createdAt, &updatedAt); err != nil {
			writeError(w, http.StatusInternalServerError, err.Error())
			return
		}

		item.ConfigData = json.RawMessage(configData)
		item.CreatedAt = createdAt.Format(time.RFC3339)
		item.UpdatedAt = updatedAt.Format(time.RFC3339)
		configs = append(configs, item)
	}

	if err := rows.Err(); err != nil {
		writeError(w, http.StatusInternalServerError, err.Error())
		return
	}

	writeJSON(w, http.StatusOK, configs)
}

func (app *App) createVisualConfig(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	var req CreateVisualConfigRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeError(w, http.StatusBadRequest, "invalid json body")
		return
	}

	if req.Name == "" {
		writeError(w, http.StatusBadRequest, "name is required")
		return
	}
	if len(req.ConfigData) == 0 || !json.Valid(req.ConfigData) {
		writeError(w, http.StatusBadRequest, "configData must be valid json")
		return
	}

	result, err := app.db.Exec(`
INSERT INTO visual_configs (name, description, config_data)
VALUES (?, ?, ?)
`, req.Name, req.Description, string(req.ConfigData))
	if err != nil {
		writeError(w, http.StatusInternalServerError, err.Error())
		return
	}

	id, err := result.LastInsertId()
	if err != nil {
		writeError(w, http.StatusInternalServerError, err.Error())
		return
	}

	writeJSON(w, http.StatusCreated, map[string]any{
		"id":      id,
		"message": "created",
	})
}

func (app *App) withCORS(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", env("CORS_ALLOW_ORIGIN", "http://localhost:5173"))
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusNoContent)
			return
		}

		next(w, r)
	}
}

func writeJSON(w http.ResponseWriter, status int, data any) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(data)
}

func writeError(w http.ResponseWriter, status int, message string) {
	writeJSON(w, status, map[string]string{"message": message})
}

func env(key string, fallback string) string {
	value := os.Getenv(key)
	if value == "" {
		return fallback
	}
	return value
}
