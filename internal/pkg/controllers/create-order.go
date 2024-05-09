package controllers

import (
	"github.com/gin-gonic/gin"
)

type CreateOrderRequest struct {
	Email string `json:"email"`
}

type CreateOrderResponse struct{}

func (s *Server) CreateOrder(c *gin.Context) {

}
