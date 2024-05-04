package controllers

import (
	"github.com/gin-gonic/gin"
)

type CreateNewsRequest struct {
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float32 `json:"price"`
}

type CreateNewsResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}

func (s *Server) CreateNews(c *gin.Context) {

}
