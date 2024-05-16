package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"oat-lab-module/internal/pkg/models"
)

type ProfileInfoRequest struct {
	Email string `json:"email"`
}

type ProfileInfoResponse struct {
	ErrorMessage string           `json:"errorMessage,omitempty"`
	ProfilesInfo []models.MedCard `json:"profileInfo"`
}

// GetProfileInfo получает информацию о профиле пользователя
//
//	@Summary		Получение информации о профиле пользователя
//	@Description	Получает информацию о профиле пользователя по email
//	@Tags			Profiles
//	@Accept			json
//	@Produce		json
//	@Param			request	body		ProfileInfoRequest	true	"Запрос для получения информации о профиле пользователя"
//	@Success		200		{object}	ProfileInfoResponse
//	@Failure		400		{object}	ProfileInfoResponse
//	@Failure		500		{object}	ProfileInfoResponse
//	@Router			/profile [get]
func (s *Server) GetProfileInfo(c *gin.Context) {
	var request ProfileInfoRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, GetCatalogResponse{ErrorMessage: err.Error()})
	}
	userProfiles, err := s.storage.GetProfiles(request.Email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, GetCatalogResponse{ErrorMessage: err.Error()})
	}

	c.JSON(http.StatusOK, ProfileInfoResponse{ProfilesInfo: userProfiles})

}
