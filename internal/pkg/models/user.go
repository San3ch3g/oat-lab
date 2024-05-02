package models

type User struct {
	Id       uint32 `gorm:"primaryKey"`
	Email    string
	Name     string
	Password string
}
