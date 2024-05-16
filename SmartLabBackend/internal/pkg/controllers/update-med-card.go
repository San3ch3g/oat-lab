package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"oat-lab-module/internal/pkg/services"
)

type UpdateMedCardRequest struct {
	MedCardId    uint32 `json:"medCardId"`
	FirstName    string `json:"firstName,omitempty"`
	LastName     string `json:"lastName,omitempty"`
	MiddleName   string `json:"middleName,omitempty"`
	BirthDate    string `json:"birthDate,omitempty"`
	Sex          string `json:"sex,omitempty"`
	MedCardImage string `json:"profileImageUrl,omitempty"`
}

type UpdateMedCardResponse struct {
	Success      bool   `json:"success"`
	ErrorMessage string `json:"errorMessage,omitempty"`
}

// UpdateMedCard Обновление информации мед карты
//
//	@Summary		Обновление информации мед карты
//	@Description	Обновляет информацию мед карты
//	@Tags			MedCards
//	@Accept			json
//	@Produce		json
//	@Param			request	body		UpdateMedCardRequest	true	"Запрос для создания профиля пользователя"
//	@Success		200		{object}	UpdateMedCardResponse
//	@Failure		400		{object}	UpdateMedCardResponse
//	@Failure		500		{object}	UpdateMedCardResponse
//	@Router			/med-card [put]
func (s *Server) UpdateMedCard(c *gin.Context) {
	var request UpdateMedCardRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, UpdateMedCardResponse{Success: false, ErrorMessage: err.Error()})
		return
	}

	if request.MedCardImage != "" {
		medCardImageUrl, err := services.SaveImage(*s.cfg, request.MedCardImage)
		if err != nil {
			c.JSON(http.StatusInternalServerError, UpdateMedCardResponse{Success: false, ErrorMessage: err.Error()})
			return
		}
		err = s.storage.UpdateMedCard(request.BirthDate, request.Sex, request.FirstName, request.MiddleName, request.LastName, medCardImageUrl, request.MedCardId)
		if err != nil {
			c.JSON(http.StatusInternalServerError, UpdateMedCardResponse{Success: false, ErrorMessage: err.Error()})
			return
		}
	} else {
		err := s.storage.UpdateMedCard(request.BirthDate, request.Sex, request.FirstName, request.MiddleName, request.LastName, "", request.MedCardId)
		if err != nil {
			c.JSON(http.StatusInternalServerError, UpdateMedCardResponse{Success: false, ErrorMessage: err.Error()})
			return
		}
	}

	c.JSON(http.StatusOK, UpdateMedCardResponse{Success: true})
}
