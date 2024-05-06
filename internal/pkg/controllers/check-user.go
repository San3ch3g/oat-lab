package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type CheckUserRequest struct {
	Email string `json:"email"`
}

type CheckUserResponse struct {
	IsRegistered bool `json:"isRegistered"`
}

func (s *Server) CheckUser(c *gin.Context) {
	var request CheckUserRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result, err := s.storage.CheckUser(*s.cfg, request.Email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	response := CheckUserResponse{
		IsRegistered: result,
	}

	c.JSON(http.StatusOK, response)
}
