package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"oat-lab-module/internal/pkg/models"
)

type GetCatalogRequest struct {
	Category string `form:"category"`
}

type GetCatalogResponse struct {
	News         []models.CatalogItem `json:"catalogItems"`
	ErrorMessage string               `json:"errorMessage,omitempty"`
}

// GetCatalogItems получает новости из каталога
//
//	@Summary		Получение новостей из каталога
//	@Description	Получает новости из каталога по категории
//	@Tags			Catalog
//	@Accept			json
//	@Produce		json
//	@Param			category	query		string	true	"Категория для получения новостей из каталога"
//	@Success		200			{object}	GetCatalogResponse
//	@Failure		400			{object}	GetCatalogResponse
//	@Failure		500			{object}	GetCatalogResponse
//	@Router			/catalog/ [get]
func (s *Server) GetCatalogItems(c *gin.Context) {
	var request GetCatalogRequest
	if err := c.ShouldBindQuery(&request); err != nil {
		c.JSON(http.StatusBadRequest, GetCatalogResponse{ErrorMessage: err.Error()})
		return
	}

	AllNews, err := s.storage.GetCatalogItems(request.Category)
	if err != nil {
		c.JSON(http.StatusInternalServerError, GetCatalogResponse{ErrorMessage: err.Error()})
		return
	}

	c.JSON(http.StatusOK, GetCatalogResponse{News: AllNews})
}
