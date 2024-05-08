package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type DeleteNewsRequest struct {
	Id uint32 `json:"id"`
}

type DeleteNewsResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}

func (s *Server) DeleteCatalogItem(c *gin.Context) {
	var request DeleteNewsRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err := s.storage.DeleteCatalogItem(request.Id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, DeleteNewsResponse{Success: true, Message: "success"})
}
