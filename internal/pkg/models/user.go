package models

type SexType string

const (
	MaleSex  SexType = "M"
	WomanSex SexType = "W"
)

type User struct {
	Id              uint32  `gorm:"primaryKey"`
	Email           string  `json:"email"`
	FirstName       string  `json:"firstName"`
	LastName        string  `json:"lastName"`
	MiddleName      string  `json:"middleName"`
	BirthDate       string  `json:"birthDate"`
	Sex             SexType `json:"sex"`
	ProfileImageUrl string  `json:"profileImageUrl"`
	Password        string
	CreatedAt       string `json:"createdAt"`
}
