package controllers

import (
	"github.com/gin-gonic/gin"
)

type DeleteProfileRequest struct {
	Id uint32 `json:"id"`
}

type DeleteProfileResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}

func (s *Server) DeleteProfile(c *gin.Context) {

}
