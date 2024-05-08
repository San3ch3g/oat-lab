package models

type SexType string

const (
	MaleSex  SexType = "M"
	WomanSex SexType = "W"
)

type User struct {
	Id        uint32 `gorm:"primaryKey"`
	Email     string `json:"email"`
	Password  string
	CreatedAt string `json:"createdAt"`
}
