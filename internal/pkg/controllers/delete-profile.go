package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type DeleteProfileRequest struct {
	Id uint32 `json:"id"`
}

type DeleteProfileResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message;omitempty"`
}

// DeleteProfile удаляет профиль пользователя
//	@Summary		Удаление профиля пользователя
//	@Description	Удаляет профиль пользователя по идентификатору
//	@Tags			Profiles
//	@Accept			json
//	@Produce		json
//	@Param			request	body		DeleteProfileRequest	true	"Запрос для удаления профиля пользователя"
//	@Success		200		{object}	DeleteProfileResponse
//	@Failure		400		{object}	DeleteProfileResponse
//	@Failure		500		{object}	DeleteProfileResponse
//	@Router			/profile [delete]
func (s *Server) DeleteProfile(c *gin.Context) {
	var request DeleteProfileRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, DeleteProfileResponse{Success: false, Message: err.Error()})
		return
	}
	err := s.storage.DeleteProfile(request.Id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, DeleteProfileResponse{Success: false, Message: err.Error()})
		return
	}
	c.JSON(http.StatusOK, DeleteProfileResponse{Success: true})
}
