package models

//api Strcuts

type User struct {
	Name          string `gorm:"column:name;"   json:"name" validate:"required"`
	Email         string `gorm:"column:email;uniqueIndex" json:"email" validate:"required,email"`
	PhoneNumber   int    `gorm:"column:phone_number" json:"phone_number" validate:"required"`
	PancardNumber string `gorm:"column:pan_card_number;uniqueIndex" json:"pan_card_number" validate:"required"`
	Password      string `gorm:"column:password" json:"password" validate:"required,min=5"`
}

func (User) TableName() string {
	return "User"
}
