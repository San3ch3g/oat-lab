package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type CheckUserRequest struct {
	Email string `json:"email"`
}

type CheckUserResponse struct {
	IsRegistered bool   `json:"isRegistered"`
	ErrorMessage string `json:"errorMessage;omitempty"`
}

// CheckUser проверяет, зарегистрирован ли пользователь
//	@Summary		Проверка пользователя
//	@Description	Проверяет, зарегистрирован ли пользователь с указанным email
//	@Tags			Auth
//	@Accept			json
//	@Produce		json
//	@Param			request	body		CheckUserRequest	true	"Запрос для проверки пользователя"
//	@Success		200		{object}	CheckUserResponse
//	@Failure		400		{object}	CheckUserResponse
//	@Failure		500		{object}	CheckUserResponse
//	@Router			/auth/check-user [post]
func (s *Server) CheckUser(c *gin.Context) {
	var request CheckUserRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, CheckUserResponse{ErrorMessage: err.Error()})
		return
	}

	result, err := s.storage.CheckUser(*s.cfg, request.Email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, CheckUserResponse{ErrorMessage: err.Error()})
	}

	response := CheckUserResponse{
		IsRegistered: result,
	}

	c.JSON(http.StatusOK, response)
}
