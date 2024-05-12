package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type DeleteCatalogItemRequest struct {
	Id uint32 `json:"id"`
}

type DeleteCatalogItemResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message;omitempty"`
}

// DeleteCatalogItem удаляет новость из каталога
//	@Summary		Удаление новости из каталога
//	@Description	Удаляет новость из каталога по идентификатору
//	@Tags			Catalog
//	@Accept			json
//	@Produce		json
//	@Param			request	body		DeleteCatalogItemRequest	true	"Запрос для удаления новости из каталога"
//	@Success		200		{object}	DeleteCatalogItemResponse
//	@Failure		400		{object}	DeleteCatalogItemResponse
//	@Failure		500		{object}	DeleteCatalogItemResponse
//	@Router			/catalog [delete]
func (s *Server) DeleteCatalogItem(c *gin.Context) {
	var request DeleteCatalogItemRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, DeleteCatalogItemResponse{Success: false, Message: err.Error()})
		return
	}
	err := s.storage.DeleteCatalogItem(request.Id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, DeleteCatalogItemResponse{Success: false, Message: err.Error()})
		return
	}
	c.JSON(http.StatusOK, DeleteCatalogItemResponse{Success: true})
}
