package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"oat-lab-module/internal/pkg/models"
)

type MedCardsInfoRequest struct {
	Email string `form:"email"`
}

type MedCardsInfoResponse struct {
	ErrorMessage string           `json:"errorMessage,omitempty"`
	MedCardsInfo []models.MedCard `json:"profileInfo"`
}

// GetMedCardsInfo получает информацию о мед картах пользователя
//
//	@Summary		Получение информации о мед картах пользователя
//	@Description	Получает информацию о мед картах пользователя по email
//	@Tags			MedCards
//	@Accept			json
//	@Produce		json
//	@Param			email	query		string	true	"Email пользователя"
//	@Success		200		{object}	MedCardsInfoResponse
//	@Failure		400		{object}	MedCardsInfoResponse
//	@Failure		500		{object}	MedCardsInfoResponse
//	@Router			/med-card [get]
func (s *Server) GetMedCardsInfo(c *gin.Context) {
	var request MedCardsInfoRequest
	if err := c.ShouldBindQuery(&request); err != nil {
		c.JSON(http.StatusBadRequest, MedCardsInfoResponse{ErrorMessage: err.Error()})
		return
	}
	userProfiles, err := s.storage.GetProfiles(request.Email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, MedCardsInfoResponse{ErrorMessage: err.Error()})
		return
	}

	c.JSON(http.StatusOK, MedCardsInfoResponse{MedCardsInfo: userProfiles})
}
