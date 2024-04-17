package models

//api Strcuts

type UserSignIn struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=5"`
}

func (UserSignIn) TableName() string {
	return "User" // Specify the actual table name here
}
