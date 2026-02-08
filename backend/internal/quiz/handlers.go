package quiz

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Handlers struct {
	Service *Service
}

func NewHandlers(s *Service) *Handlers {
	return &Handlers{Service: s}
}

func (h *Handlers) Start(c *gin.Context) {
	userID := c.GetInt("user_id")
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "5"))

	qs, err := h.Service.Start(userID, limit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, qs)
}

func (h *Handlers) Answer(c *gin.Context) {
	userID := c.GetInt("user_id")
	var req struct {
		Correct bool `json:"correct"`
	}
	c.BindJSON(&req)

	nextID, err := h.Service.Answer(userID, req.Correct)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"next_question_id": nextID})
}
