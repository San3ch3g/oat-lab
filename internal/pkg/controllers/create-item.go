package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type CreateCatalogRequest struct {
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float32 `json:"price"`
	Category    string  `json:"category,omitempty"`
}

type CreateCatalogResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message,omitempty"`
}

// CreateCatalogItem создает новую новость в каталоге
//
//	@Summary		Создание новости
//	@Description	Создает новую новость в каталоге
//	@Tags			Catalog
//	@Accept			json
//	@Produce		json
//	@Param			request	body		CreateCatalogRequest	true	"Запрос для создания новости"
//	@Success		201		{object}	CreateCatalogResponse
//	@Failure		400		{object}	CreateCatalogResponse
//	@Failure		500		{object}	CreateCatalogResponse
//	@Router			/catalog [post]
func (s *Server) CreateCatalogItem(c *gin.Context) {
	var request CreateCatalogRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, CreateCatalogResponse{Success: false, Message: err.Error()})
		return
	}
	err := s.storage.CreateCatalogItem(request.Name, request.Description, request.Category, request.Price)
	if err != nil {
		c.JSON(http.StatusInternalServerError, CreateCatalogResponse{Success: false, Message: err.Error()})
		return
	}
	c.JSON(http.StatusCreated, CreateCatalogResponse{Success: true})
}
