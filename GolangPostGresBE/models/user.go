package models

type User struct {
	Id          uint   `json:"id"`
	FirstName   string `json:"firstName"`
	LastName    string `json:"lastName"`
	PhoneNumber int    `json:"phoneNumber"`
	Email       string `json:"email" gorm:"unique"`
	UserType    string `json:"userType"`
	Sou         bool   `json:"sou"`
	Password    []byte `json:"-"`
}
