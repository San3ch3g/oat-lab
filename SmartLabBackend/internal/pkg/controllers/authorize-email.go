package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type AuthorizeRequest struct {
	Email string `json:"email"`
}
type AuthorizeResponse struct {
	Success      bool   `json:"success"`
	ErrorMessage string `json:"errorMessage,omitempty"`
	Id           uint32 `json:"id,omitempty"`
}

// Authorize godoc
//
//	@Summary		Проверка email
//	@Description	Отправляет код на email и создаёт запись в бд
//	@Tags			Auth
//	@Accept			json
//	@Produce		json
//	@Param			request	body		AuthorizeRequest	true	"Запрос для отправки кода и создания записи в бд"
//	@Success		200		{object}	AuthorizeResponse
//	@Failure		400		{object}	AuthorizeResponse
//	@Failure		500		{object}	AuthorizeResponse
//	@Router			/auth/authorize [post]
func (s *Server) Authorize(c *gin.Context) {
	var request AuthorizeRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, AuthorizeResponse{Success: false, ErrorMessage: err.Error()})
		return
	}
	userID, err := s.storage.Authorize(*s.cfg, request.Email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, AuthorizeResponse{Success: false, ErrorMessage: err.Error()})
		return
	}
	c.JSON(http.StatusOK, AuthorizeResponse{Success: true, Id: userID})
}
