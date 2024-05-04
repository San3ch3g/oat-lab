package controllers

import (
	"github.com/gin-gonic/gin"
)

type SignInRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type SignInResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}

func (s *Server) SignIn(c *gin.Context) {

}
