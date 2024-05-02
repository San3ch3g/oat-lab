package models

type SexType string

const (
	MaleSex  SexType = "M"
	WomanSex SexType = "W"
)

type User struct {
	Id              uint32 `gorm:"primaryKey"`
	Email           string
	FirstName       string
	LastName        string
	MiddleName      string
	BirthDate       string
	Sex             SexType
	ProfileImageUrl string
	Password        string
}
