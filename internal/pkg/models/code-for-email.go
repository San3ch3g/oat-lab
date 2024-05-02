package models

type CodeForEmail struct {
	Id        uint   `gorm:"primaryKey"`
	Email     string `gorm:"unique"`
	Code      string
	CreatedAt int64
}
