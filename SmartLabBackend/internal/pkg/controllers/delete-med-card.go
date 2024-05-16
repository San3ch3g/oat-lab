package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type MedCardRequest struct {
	Id uint32 `json:"id"`
}

type MedCardResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message,omitempty"`
}

// DeleteMedCard удаляет профиль пользователя
//
//	@Summary		Удаление профиля пользователя
//	@Description	Удаляет профиль пользователя по идентификатору
//	@Tags			MedCards
//	@Accept			json
//	@Produce		json
//	@Param			request	body		MedCardRequest	true	"Запрос для удаления профиля пользователя"
//	@Success		200		{object}	MedCardResponse
//	@Failure		400		{object}	MedCardResponse
//	@Failure		500		{object}	MedCardResponse
//	@Router			/med-card [delete]
func (s *Server) DeleteMedCard(c *gin.Context) {
	var request MedCardRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, MedCardResponse{Success: false, Message: err.Error()})
		return
	}
	err := s.storage.DeleteProfile(request.Id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, MedCardResponse{Success: false, Message: err.Error()})
		return
	}
	c.JSON(http.StatusOK, MedCardResponse{Success: true})
}
