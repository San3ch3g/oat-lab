package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type CreateNewsRequest struct {
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float32 `json:"price"`
	Category    string  `json:"category,omitempty"`
}

type CreateNewsResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}

func (s *Server) CreateNews(c *gin.Context) {
	var request CreateNewsRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err := s.storage.CreateNews(request.Name, request.Description, request.Category, request.Price)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, CreateNewsResponse{Success: true, Message: "News created successfully!"})
}
