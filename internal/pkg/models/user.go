package models

type User struct {
	Id        uint32 `gorm:"primaryKey"`
	Email     string `json:"email"`
	CreatedAt string `json:"createdAt"`
}
