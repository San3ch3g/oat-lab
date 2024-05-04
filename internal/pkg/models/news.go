package models

type News struct {
	Id          uint32  `gorm:"primaryKey" json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float32 `json:"price"`
	ImageUrl    string  `json:"imageUrl"`
	CreatedAt   string  `json:"createdAt"`
}
