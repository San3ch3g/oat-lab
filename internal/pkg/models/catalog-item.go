package models

type CatalogItemCategory string

const (
	Default CatalogItemCategory = "all"
	News    CatalogItemCategory = "news"
	Covid   CatalogItemCategory = "covid"
	Popular CatalogItemCategory = "popular"
	Complex CatalogItemCategory = "complex"
)

type CatalogItem struct {
	Id          uint32              `gorm:"primaryKey" json:"id"`
	Name        string              `json:"name"`
	Description string              `json:"description"`
	Price       float32             `json:"price"`
	Category    CatalogItemCategory `json:"category" gorm:"default:'all'"`
	ImageUrl    string              `json:"imageUrl"`
	CreatedAt   string              `json:"createdAt"`
}
