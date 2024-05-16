package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"oat-lab-module/internal/pkg/models"
)

type CreateMedCardRequest struct {
	Email        string         `json:"email"`
	FirstName    string         `json:"firstName"`
	LastName     string         `json:"lastName"`
	MiddleName   string         `json:"middleName"`
	BirthDate    string         `json:"birthDate"`
	ProfileImage string         `json:"profileImage,omitempty"`
	Sex          models.SexType `json:"sex"`
}

type CreateMedCardResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message,omitempty"`
}

// CreateMedCard создает профиль пользователя
//
//	@Summary		Создание профиля пользователя
//	@Description	Создает профиль пользователя с указанными данными
//	@Tags			MedCards
//	@Accept			json
//	@Produce		json
//	@Param			request	body		CreateMedCardRequest	true	"Запрос для создания профиля пользователя"
//	@Success		201		{object}	CreateMedCardResponse
//	@Failure		400		{object}	CreateMedCardResponse
//	@Failure		500		{object}	CreateMedCardResponse
//	@Router			/med-card [post]
func (s *Server) CreateMedCard(c *gin.Context) {
	var request CreateMedCardRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, CreateMedCardResponse{Success: false, Message: err.Error()})
		return
	}
	err := s.storage.CreateProfile(*s.cfg, request.Email, request.FirstName, request.LastName, request.MiddleName, request.BirthDate, request.Sex, request.ProfileImage)
	if err != nil {
		c.JSON(http.StatusInternalServerError, CreateMedCardResponse{Success: false, Message: err.Error()})
		return
	}
	c.JSON(http.StatusCreated, CreateMedCardResponse{Success: true})
}
