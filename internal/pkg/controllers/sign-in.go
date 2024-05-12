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

// SignIn выполняет вход пользователя в систему
//	@Summary		Вход пользователя
//	@Description	Выполняет вход пользователя в систему с использованием email и пароля
//	@Tags			Auth
//	@Accept			json
//	@Produce		json
//	@Param			request	body		SignInRequest	true	"Запрос для входа пользователя"
//	@Success		200		{object}	SignInResponse
//	@Failure		400		{object}	SignInResponse
//	@Failure		401		{object}	SignInResponse
//	@Router			/auth/sign-in [post]
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
