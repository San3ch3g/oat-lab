package models

type News struct {
	Id          uint32 `gorm:"primaryKey"`
	Name        string
	Description string
	Price       uint32
	ImageUrl    string
	CreatedAt   int64
}
