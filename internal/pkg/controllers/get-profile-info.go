package controllers

import (
	"github.com/gin-gonic/gin"
	"oat-lab-module/internal/pkg/models"
)

type ProfileInfoRequest struct {
	Email string `json:"email"`
}

type ProfileInfoResponse struct {
	ProfileInfo models.User `json:"profileInfo"`
}

func (s *Server) GetProfileInfo(c *gin.Context) {

}
