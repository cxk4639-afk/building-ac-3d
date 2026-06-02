package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	App      AppConfig
	Database DatabaseConfig
	Auth     AuthConfig
	CORS     CORSConfig
}

type AppConfig struct {
	Port string
}

type DatabaseConfig struct {
	Host     string
	Port     string
	Name     string
	User     string
	Password string
}

type AuthConfig struct {
	JWTSecret string
}

type CORSConfig struct {
	AllowOrigin string
}

func Load() *Config {
	_ = godotenv.Load()
	_ = godotenv.Load("backend/.env")

	return &Config{
		App: AppConfig{Port: env("APP_PORT", "8080")},
		Database: DatabaseConfig{
			Host:     env("MYSQL_HOST", "localhost"),
			Port:     env("MYSQL_PORT", "3306"),
			Name:     env("MYSQL_DATABASE", "building_ac_3d"),
			User:     env("MYSQL_USER", "root"),
			Password: env("MYSQL_PASSWORD", ""),
		},
		Auth: AuthConfig{JWTSecret: env("JWT_SECRET", "building-ac-3d-dev-secret")},
		CORS: CORSConfig{AllowOrigin: env("CORS_ALLOW_ORIGIN", "http://localhost:5173")},
	}
}

func env(key string, fallback string) string {
	value := os.Getenv(key)
	if value == "" {
		return fallback
	}
	return value
}
