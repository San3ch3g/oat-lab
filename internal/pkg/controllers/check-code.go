package controllers

import (
	"github.com/gin-gonic/gin"
)

type CheckCodeRequest struct {
	Email string `json:"email"`
	Code  string `json:"code"`
}

type CheckCodeResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}

func (s *Server) CheckCode(c *gin.Context) {

}
