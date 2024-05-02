package models

type Order struct {
	Id    int32 `gorm:"primaryKey"`
	Name  string
	Price float32
}
