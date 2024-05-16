package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

type GetCatalogItemByIdRequest struct {
	Id uint32 `form:"itemId"`
}

type GetCatalogItemByIdResponse struct {
	ErrorMessage string `json:"errorMessage,omitempty"`
	Id           uint32 `gorm:"primaryKey" json:"id"`
	Name         string `json:"name"`
	Description  string `json:"description"`
	Price        string `json:"price"`
	TimeRes      string `json:"timeRes"`
	Preparation  string `json:"preparation"`
	BIO          string `json:"bio"`
	Category     string `json:"category"`
	ImageUrl     string `json:"imageUrl"`
	CreatedAt    string `json:"createdAt"`
}

// GetCatalogItemById godoc
// @Summary Получить элемент каталога по ID.
// @Description Получает элемент каталога по его ID.
// @Tags Catalog
// @Accept json
// @Produce json
// @Param itemId query uint32 true "ID элемента каталога"
// @Success 200 {object} GetCatalogItemByIdResponse
// @Failure 400 {object} GetCatalogItemByIdResponse
// @Failure 500 {object} GetCatalogItemByIdResponse
// @Router /catalog/by-id [get]
func (s *Server) GetCatalogItemById(c *gin.Context) {
	var request GetCatalogItemByIdRequest
	if err := c.ShouldBindQuery(&request); err != nil {
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
		TimeRes:     catalogItem.TimeRes,
		Preparation: catalogItem.Preparation,
		BIO:         catalogItem.BIO,
		Category:    fmt.Sprintf("%v", catalogItem.Category),
		CreatedAt:   catalogItem.CreatedAt,
	})
}
