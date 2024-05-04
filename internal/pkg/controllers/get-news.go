package controllers

import (
	"github.com/gin-gonic/gin"
	"oat-lab-module/internal/pkg/models"
)

type GetNewsRequest struct{}

type PostNewsRequest struct {
	News []models.News `json:"news"`
}

func (s *Server) GetNews(c *gin.Context) {

}
