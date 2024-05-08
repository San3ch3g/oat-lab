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
	News []models.CatalogItem `json:"catalogItems"`
}

func (s *Server) GetCatalogItems(c *gin.Context) {
	var request GetNewsRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	AllNews, err := s.storage.GetCatalogItems(request.Category)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	c.JSON(http.StatusOK, GetNewsResponse{AllNews})

}
