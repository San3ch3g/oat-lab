package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"oat-lab-module/internal/pkg/models"
)

type ProfileInfoRequest struct {
	Email string `json:"email"`
}

type ProfileInfoResponse struct {
	ProfileInfo []models.Profile `json:"profileInfo"`
}

func (s *Server) GetProfileInfo(c *gin.Context) {
	var request ProfileInfoRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	userProfiles, err := s.storage.GetProfiles(request.Email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
	}

	c.JSON(http.StatusOK, ProfileInfoResponse{userProfiles})

}
