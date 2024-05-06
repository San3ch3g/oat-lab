package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type SendCodeAgainRequest struct {
	Email string `json:"email"`
}

type SendCodeAgainResponse struct {
	Success bool `json:"success"`
}

func (s *Server) SendCodeAgain(c *gin.Context) {
	var request SendCodeAgainRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, SendCodeAgainResponse{
			Success: false,
		})
		return
	}
	err := s.storage.SendCodeAgain(*s.cfg, request.Email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, nil)
	}
}
