package models

type CodeForEmail struct {
	Id        uint32 `gorm:"primaryKey"`
	Email     string `gorm:"unique"`
	Code      string
	CreatedAt string
}
