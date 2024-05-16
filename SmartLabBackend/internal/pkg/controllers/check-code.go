package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type CheckCodeRequest struct {
	Email string `json:"email"`
	Code  string `json:"code"`
}

type CheckCodeResponse struct {
	UserId  uint32 `json:"userId"`
	Message string `json:"message,omitempty"`
}

// CheckCode godoc
//
//	@Summary		Проверка кода
//	@Description	Проверяет код, отправленный на указанный email
//	@Tags			Auth
//	@Accept			json
//	@Produce		json
//	@Param			request	body		CheckCodeRequest	true	"Запрос на проверку кода"
//	@Success		200		{object}	CheckCodeResponse
//	@Failure		400		{object}	CheckCodeResponse
//	@Failure		422		{object}	CheckCodeResponse
//	@Router			/auth/check-code [post]
func (s *Server) CheckCode(c *gin.Context) {
	var request CheckCodeRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, CheckCodeResponse{
			Message: err.Error(),
		})
		return
	}

	userId, err := s.storage.CheckCode(request.Email, request.Code)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, CheckCodeResponse{
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, CheckCodeResponse{
		UserId: userId,
	})
}
