package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"oat-lab-module/internal/pkg/models"
)

type FillProfileRequest struct {
	Email        string         `json:"email"`
	FirstName    string         `json:"firstName"`
	LastName     string         `json:"lastName"`
	MiddleName   string         `json:"middleName"`
	BirthDate    string         `json:"birthDate"`
	ProfileImage []byte         `json:"profileImage,omitempty"`
	Sex          models.SexType `json:"sex"`
}

type FillProfileResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}

func (s *Server) FillProfile(c *gin.Context) {
	var request FillProfileRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err := s.storage.FillProfile(request.Email, request.FirstName, request.LastName, request.MiddleName, request.BirthDate, request.Sex)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
	c.JSON(http.StatusOK, FillProfileResponse{Success: true, Message: "ok"})
}
