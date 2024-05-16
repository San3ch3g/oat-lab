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

// SendCodeAgain повторно отправляет код подтверждения на email пользователя
//	@Summary		Повторная отправка кода подтверждения
//	@Description	Повторно отправляет код подтверждения на email пользователя
//	@Tags			Auth
//	@Accept			json
//	@Produce		json
//	@Param			request	body		SendCodeAgainRequest	true	"Запрос для повторной отправки кода подтверждения"
//	@Success		200		{object}	SendCodeAgainResponse
//	@Failure		400		{object}	SendCodeAgainResponse
//	@Failure		500		{object}	SendCodeAgainResponse
//	@Router			/auth/send-code-again [post]
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
