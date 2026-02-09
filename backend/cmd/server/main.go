package main

import (
	"log"
	"net/http"
	"os"

	"interview-copilot/backend/internal/config"
	"interview-copilot/backend/internal/db"
	"interview-copilot/backend/internal/http/router"
	"interview-copilot/backend/internal/metrics"

	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {
	cfg := config.Load()

	db.Connect(cfg.DatabaseURL)
	metrics.Init()

	r := router.New()

	r.GET("/metrics", ginWrap(promhttp.Handler()))
	r.GET("/health", ginWrap(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		w.WriteHeader(200)
		w.Write([]byte("ok"))
	})))

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Println("Backend running on :" + port)
	r.Run(":" + port)
}

func ginWrap(h http.Handler) func(*gin.Context) {
	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}
