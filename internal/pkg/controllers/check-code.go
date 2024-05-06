package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
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
	var request CheckCodeRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, CheckCodeResponse{
			Success: false,
			Message: err.Error(),
		})
		return
	}

	err := s.storage.CheckCode(request.Email, request.Code)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, CheckCodeResponse{
			Success: false,
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, CheckCodeResponse{
		Success: true,
		Message: "Code is valid",
	})
}
