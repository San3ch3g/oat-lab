package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type SignUpRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type SignUpResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}

// SignUp регистрирует нового пользователя
//	@Summary		Регистрация пользователя
//	@Description	Регистрирует нового пользователя с использованием email и пароля
//	@Tags			Auth
//	@Accept			json
//	@Produce		json
//	@Param			request	body		SignUpRequest	true	"Запрос для регистрации пользователя"
//	@Success		200		{object}	SignUpResponse
//	@Failure		400		{object}	SignUpResponse
//	@Failure		500		{object}	SignUpResponse
//	@Router			/auth/sign-up [post]
func (s *Server) SignUp(c *gin.Context) {
	var request SignUpRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, SignUpResponse{
			Success: false,
			Message: "Invalid request body",
		})
		return
	}

	err := s.storage.SignUp(request.Email, request.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, SignUpResponse{
			Success: false,
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, SignUpResponse{
		Success: true,
		Message: "User created successfully",
	})
}
