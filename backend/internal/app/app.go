package app

import (
	"log"

	"interview-copilot/backend/internal/config"
	"interview-copilot/backend/internal/db"
	"interview-copilot/backend/internal/http"
)

func Run() {
	cfg := config.Load()

	database := db.Connect(cfg.DatabaseURL)
	db.Migrate(database)

	router := http.NewRouter(database, cfg)

	log.Println("Backend listening on :8080")
	router.Run(":8080")
}
