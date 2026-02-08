package http

import (
	"interview-copilot/backend/internal/auth"
	"interview-copilot/backend/internal/config"
	"interview-copilot/backend/internal/http/handlers"
	"interview-copilot/backend/internal/http/middleware"
	"interview-copilot/backend/internal/repository"
	"interview-copilot/backend/internal/service"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
)

func NewRouter(db *pgxpool.Pool, cfg *config.Config) *gin.Engine {
	r := gin.New()
	r.Use(gin.Recovery(), middleware.Logger())

	userRepo := &repository.UserRepo{DB: db}
	authSvc := &service.AuthService{Users: userRepo, JWTSecret: cfg.JWTSecret}
	h := handlers.New(authSvc)

	r.GET("/health", h.Health)
	r.POST("/auth/register", h.Register)
	r.POST("/auth/login", h.Login)

	api := r.Group("/api")
	api.Use(auth.Middleware(cfg.JWTSecret))
	{
		api.GET("/me", func(c *gin.Context) {
			c.JSON(200, gin.H{"user_id": c.GetInt("user_id")})
		})
	}

	return r
}
