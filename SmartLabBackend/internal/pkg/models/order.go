package models

type Order struct {
	Id      uint32 `gorm:"primaryKey"`
	Address string `json:"address"`
	Date    int64  `json:"date"`
	Phone   string `json:"phone"`
	Comment string `json:"comment"`
}
