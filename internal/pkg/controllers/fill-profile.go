package controllers

import (
	"github.com/gin-gonic/gin"
	"oat-lab-module/internal/pkg/models"
)

type FillProfileRequest struct {
	FirstName    string         `json:"firstName"`
	LastName     string         `json:"lastName"`
	MiddleName   string         `json:"middleName"`
	BirthDate    string         `json:"birthDate"`
	ProfileImage []byte         `json:"profileImage"`
	Sex          models.SexType `json:"sex"`
}

type FillProfileResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}

func (s *Server) FillProfile(c *gin.Context) {

}
