package progress

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Handlers struct {
	Repo *Repository
}

func NewHandlers(r *Repository) *Handlers {
	return &Handlers{Repo: r}
}

func (h *Handlers) Stats(c *gin.Context) {
	userID := c.GetInt("user_id")
	total, correct, err := h.Repo.Stats(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"total":   total,
		"correct": correct,
		"ratio":   float64(correct) / float64(max(total, 1)),
	})
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
