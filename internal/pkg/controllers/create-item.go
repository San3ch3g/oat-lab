package controllers

import (
	"github.com/gin-gonic/gin"
)

type CreateItemRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Category    string `json:"category"`
	Price       uint32 `json:"price"`
}

type CreateItemResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}

func (s *Server) CreateItem(c *gin.Context) {

}
