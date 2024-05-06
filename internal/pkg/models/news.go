package models

type NewsCategory string

const (
	Default NewsCategory = "all"
	Covid   NewsCategory = "covid"
	Popular NewsCategory = "popular"
	Complex NewsCategory = "complex"
)

type News struct {
	Id          uint32       `gorm:"primaryKey" json:"id"`
	Name        string       `json:"name"`
	Description string       `json:"description"`
	Price       float32      `json:"price"`
	Category    NewsCategory `json:"category" gorm:"default:'all'"`
	ImageUrl    string       `json:"imageUrl"`
	CreatedAt   string       `json:"createdAt"`
}
