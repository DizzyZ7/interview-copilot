package config

import "os"

type Config struct {
	DatabaseURL string
	JWTSecret   string
	Env         string
}

func Load() *Config {
	return &Config{
		DatabaseURL: getenv("DATABASE_URL", "postgres://copilot:copilot@db:5432/copilot?sslmode=disable"),
		JWTSecret:   getenv("JWT_SECRET", "supersecret"),
		Env:         getenv("ENV", "dev"),
	}
}

func getenv(k, def string) string {
	if v := os.Getenv(k); v != "" {
		return v
	}
	return def
}
