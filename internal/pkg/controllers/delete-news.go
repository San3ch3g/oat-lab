package controllers

import "github.com/gin-gonic/gin"

type DeleteNewsRequest struct {
	Id uint32 `json:"id"`
}

type DeleteNewsResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}

func (s *Server) DeleteNews(c *gin.Context) {

}
