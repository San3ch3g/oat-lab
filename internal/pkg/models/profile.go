package models

type Profile struct {
	Id              uint32  `json:"id" gorm:"primaryKey"`
	UserId          uint32  `json:"userId"`
	FirstName       string  `json:"firstName"`
	LastName        string  `json:"lastName"`
	MiddleName      string  `json:"middleName"`
	BirthDate       string  `json:"birthDate"`
	Sex             SexType `json:"sex" gorm:"type:sex_type"`
	ProfileImageUrl string  `json:"profileImageUrl"`
}
