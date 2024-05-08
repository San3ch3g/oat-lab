package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type DeleteProfileRequest struct {
	Id uint32 `json:"id"`
}

type DeleteProfileResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}

func (s *Server) DeleteProfile(c *gin.Context) {
	var request DeleteProfileRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err := s.storage.DeleteProfile(request.Id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, nil)
		return
	}
	c.JSON(http.StatusOK, DeleteProfileResponse{Success: true, Message: "Profile deleted successfully!"})
}
