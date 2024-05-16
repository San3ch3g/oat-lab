package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"oat-lab-module/internal/pkg/services"
)

type SendCodeRequest struct {
	Email string `json:"email"`
}
type SendCodeResponse struct {
	Success      bool   `json:"success"`
	ErrorMessage string `json:"errorMessage,omitempty"`
}

// SendCode godoc
//
//	@Summary		Проверка email
//	@Description	Отправляет код на email и создаёт запись в бд
//	@Tags			Auth
//	@Accept			json
//	@Produce		json
//	@Param			request	body		SendCodeRequest	true	"Запрос для отправки кода и создания записи в бд"
//	@Success		200		{object}	SendCodeResponse
//	@Failure		400		{object}	SendCodeResponse
//	@Failure		500		{object}	SendCodeResponse
//	@Router			/auth/send-code [post]
func (s *Server) SendCode(c *gin.Context) {
	var request SendCodeRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, SendCodeResponse{Success: false, ErrorMessage: err.Error()})
		return
	}
	isValid := services.ValidEmail(request.Email)
	if !isValid {
		c.JSON(http.StatusBadRequest, SendCodeResponse{ErrorMessage: "invalid email"})
		return
	}
	err := s.storage.SendCode(*s.cfg, request.Email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, SendCodeResponse{Success: false, ErrorMessage: err.Error()})
		return
	}
	c.JSON(http.StatusOK, SendCodeResponse{Success: true})
}
