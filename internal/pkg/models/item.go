package models

type Item struct {
	Id          uint32 `gorm:"primaryKey" json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Category    string `json:"category"`
	Price       uint32 `json:"price"`
	CreatedAt   string `json:"createdAt"`
}
