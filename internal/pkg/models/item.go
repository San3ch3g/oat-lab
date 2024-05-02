package models

type Item struct {
	Id          uint32 `gorm:"primaryKey"`
	Name        string
	Description string
	Category    string
	Price       uint32
}
