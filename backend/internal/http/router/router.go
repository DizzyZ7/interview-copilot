package router

import (
	"net/http"

	"interview-copilot/backend/internal/auth"
	"interview-copilot/backend/internal/config"
	"interview-copilot/backend/internal/health"
	"interview-copilot/backend/internal/http/handlers"
	"interview-copilot/backend/internal/http/middleware"
	"interview-copilot/backend/internal/metrics"
	"interview-copilot/backend/internal/progress"
	"interview-copilot/backend/internal/questions"
	"interview-copilot/backend/internal/quiz"
	"interview-copilot/backend/internal/repository"
	"interview-copilot/backend/internal/service"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func NewRouter(db *pgxpool.Pool, cfg *config.Config) *gin.Engine {
	r := gin.New()

	// Core middleware
	r.Use(gin.Recovery())
	r.Use(middleware.Logger())
	r.Use(metrics.Middleware())

	// Infra
	r.GET("/health", health.Handler)
	r.GET("/metrics", gin.WrapH(promhttp.Handler()))

	// Auth
	userRepo := &repository.UserRepo{DB: db}
	authSvc := &service.AuthService{Users: userRepo, JWTSecret: cfg.JWTSecret}
	authHandlers := handlers.New(authSvc)

	r.POST("/auth/register", authHandlers.Register)
	r.POST("/auth/login", authHandlers.Login)

	// Domains
	qRepo := &questions.Repository{DB: db}
	qSvc := &questions.Service{Repo: qRepo}
	qHandlers := questions.NewHandlers(qSvc)

	quizSvc := quiz.NewService(qSvc)
	quizHandlers := quiz.NewHandlers(quizSvc)

	progressRepo := &progress.Repository{DB: db}
	progressHandlers := progress.NewHandlers(progressRepo)

	// Versioned API
	api := r.Group("/api")
	api.Use(auth.Middleware(cfg.JWTSecret))
	{
		// Questions
		api.POST("/questions", qHandlers.Create)
		api.GET("/questions", qHandlers.List)
		api.GET("/questions/:id", qHandlers.Get)
		api.PUT("/questions/:id", qHandlers.Update)
		api.DELETE("/questions/:id", qHandlers.Delete)

		// Random / Quiz
		api.GET("/random", qHandlers.Random)
		api.POST("/quiz/start", quizHandlers.Start)
		api.POST("/quiz/answer", quizHandlers.Answer)

		// Progress
		api.GET("/progress", progressHandlers.Stats)

		// Me
		api.GET("/me", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"user_id": c.GetInt("user_id")})
		})
	}

	// 404 fallback
	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{"error": "route not found"})
	})

	return r
}
