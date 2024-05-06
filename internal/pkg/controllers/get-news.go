package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"oat-lab-module/internal/pkg/models"
)

type GetNewsRequest struct {
	Category string `json:"category"`
}

type GetNewsResponse struct {
	News []models.News `json:"news"`
}

func (s *Server) GetNews(c *gin.Context) {
	var request GetNewsRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	AllNews, err := s.storage.GetNews(request.Category)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	c.JSON(http.StatusOK, GetNewsResponse{AllNews})

}
