package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
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
	var request SignInRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, SignInResponse{
			Success: false,
			Message: "Invalid request body",
		})
		return
	}

	err := s.storage.SignIn(request.Email, request.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, SignInResponse{
			Success: false,
			Message: "Invalid email or password",
		})
		return
	}

	c.JSON(http.StatusOK, SignInResponse{
		Success: true,
		Message: "User signed in successfully",
	})
}
