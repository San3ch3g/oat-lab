package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

type GetCatalogItemByIdRequest struct {
	Id uint32 `json:"itemId"`
}
type GetCatalogItemByIdResponse struct {
	ErrorMessage string  `json:"errorMessage,omitempty"`
	Id           uint32  `gorm:"primaryKey" json:"id"`
	Name         string  `json:"name"`
	Description  string  `json:"description"`
	Price        float32 `json:"price"`
	Category     string  `json:"category"`
	ImageUrl     string  `json:"imageUrl"`
	CreatedAt    string  `json:"createdAt"`
}

// GetCatalogItemById godoc
// @Summary Получить элемент каталога по ID.
// @Description Получает элемент каталога по его ID.
// @Tags Catalog
// @Accept json
// @Produce json
// @Param request body GetCatalogItemByIdRequest true "Тело запроса с itemId"
// @Success 200 {object} GetCatalogItemByIdResponse
// @Failure 400 {object} GetCatalogItemByIdResponse
// @Failure 500 {object} GetCatalogItemByIdResponse
// @Router /catalog/get-by-id [post]
func (s *Server) GetCatalogItemById(c *gin.Context) {
	var request GetCatalogItemByIdRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(400, GetCatalogItemByIdResponse{ErrorMessage: fmt.Sprintf("%v", err)})
		return
	}
	catalogItem, err := s.storage.GetCatalogItemById(request.Id)
	if err != nil {
		c.JSON(500, GetCatalogItemByIdResponse{ErrorMessage: fmt.Sprintf("%v", err)})
		return
	}
	c.JSON(200, GetCatalogItemByIdResponse{
		Id:          catalogItem.Id,
		Name:        catalogItem.Name,
		Description: catalogItem.Description,
		Price:       catalogItem.Price,
		Category:    fmt.Sprintf("%v", catalogItem.Category),
		ImageUrl:    catalogItem.ImageUrl,
		CreatedAt:   catalogItem.CreatedAt,
	})
}
