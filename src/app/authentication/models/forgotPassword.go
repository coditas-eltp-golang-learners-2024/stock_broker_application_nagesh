package models

//api Strcuts

type ForgotPassword struct {
	Email         string `json:"email" validate:"required,email"`
	PanCardNumber string `json:"pan_card_number" validate:"required"`
	NewPassword   string `json:"new_password" validate:"required,passwordValidation,min=5"`
}

func (ForgotPassword) TableName() string {
	return "User" // Specify the actual table name here
}
