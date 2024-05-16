package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"oat-lab-module/internal/pkg/models"
)

type CreateMedCardRequest struct {
	Email      string         `json:"email"`
	FirstName  string         `json:"firstName"`
	LastName   string         `json:"lastName"`
	MiddleName string         `json:"middleName"`
	BirthDate  string         `json:"birthDate"`
	Sex        models.SexType `json:"sex"`
}

type CreateMedCardResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message,omitempty"`
}

// CreateMedCard создает мед карту пользователя
//
//	@Summary		Создание мед карты пользователя
//	@Description	Создает мед карты пользователя с указанными данными
//	@Tags			MedCards
//	@Accept			json
//	@Produce		json
//	@Param			request	body		CreateMedCardRequest	true	"Запрос для создания мед карты пользователя"
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
	fmt.Printf("request Email = '%v'", request.Email)
	err := s.storage.CreateMedCard(*s.cfg, request.Email, request.FirstName, request.LastName, request.MiddleName, request.BirthDate, request.Sex)
	if err != nil {
		c.JSON(http.StatusInternalServerError, CreateMedCardResponse{Success: false, Message: err.Error()})
		return
	}
	c.JSON(http.StatusCreated, CreateMedCardResponse{Success: true})
}
