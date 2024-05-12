package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"oat-lab-module/internal/pkg/models"
)

type CreateProfileRequest struct {
	Email        string         `json:"email"`
	FirstName    string         `json:"firstName"`
	LastName     string         `json:"lastName"`
	MiddleName   string         `json:"middleName"`
	BirthDate    string         `json:"birthDate"`
	ProfileImage []byte         `json:"profileImage,omitempty"`
	Sex          models.SexType `json:"sex"`
}

type CreateProfileResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message;omitempty"`
}

// CreateProfile создает профиль пользователя
//	@Summary		Создание профиля пользователя
//	@Description	Создает профиль пользователя с указанными данными
//	@Tags			Profiles
//	@Accept			json
//	@Produce		json
//	@Param			request	body		CreateProfileRequest	true	"Запрос для создания профиля пользователя"
//	@Success		201		{object}	CreateProfileResponse
//	@Failure		400		{object}	CreateProfileResponse
//	@Failure		500		{object}	CreateProfileResponse
//	@Router			/profile [post]
func (s *Server) CreateProfile(c *gin.Context) {
	var request CreateProfileRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, CreateProfileResponse{Success: false, Message: err.Error()})
		return
	}
	err := s.storage.CreateProfile(*s.cfg, request.Email, request.FirstName, request.LastName, request.MiddleName, request.BirthDate, request.Sex, request.ProfileImage)
	if err != nil {
		c.JSON(http.StatusInternalServerError, CreateProfileResponse{Success: false, Message: err.Error()})
		return
	}
	c.JSON(http.StatusCreated, CreateProfileResponse{Success: true})
}
