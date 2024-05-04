package controllers

import (
	"github.com/gin-gonic/gin"
	"oat-lab-module/internal/pkg/models"
)

type GetItemsRequest struct{}

type GetItemsResponse struct {
	Items []models.Item `json:"items"`
}

func (s *Server) GetItems(c *gin.Context) {

}
