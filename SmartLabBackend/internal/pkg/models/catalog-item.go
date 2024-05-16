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
	Price       string              `json:"price"`
	TimeRes     string              `json:"timeResult"`
	Preparation string              `json:"preparation"`
	BIO         string              `json:"bio"`
	Category    CatalogItemCategory `json:"category" gorm:"default:'all'"`
	CreatedAt   string              `json:"createdAt"`
}
