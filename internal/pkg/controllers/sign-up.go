package controllers

import (
	"github.com/gin-gonic/gin"
)

type SignUpRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type SignUpResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}

func (s *Server) SignUp(c *gin.Context) {

}
