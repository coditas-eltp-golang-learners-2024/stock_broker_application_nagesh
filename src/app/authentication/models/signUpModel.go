package models

//api Strcuts

type User struct {
	Name          string `json:"name" validate:"required"`
	Email         string `json:"email" validate:"required,email"`
	PhoneNumber   int    `json:"phone_number" validate:"required"`
	PancardNumber string `json:"pan_card_number" validate:"required"`
	Password      string `json:"password" validate:"required,min=5"`
}
