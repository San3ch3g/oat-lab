package controllers

import (
	"github.com/gin-gonic/gin"
)

type CheckUserRequest struct {
	Email string `json:"email"`
}

type CheckUserResponse struct {
	IsRegistered bool `json:"isRegistered"`
}

func (s *Server) CheckUser(c *gin.Context) {

}
