package models

import "time"

type ValidateOtp struct {
	Otp    int       `json:"otp" validate:"required"`
	OtpExp time.Time `json:"otpExp"`
	Email  string    `json:"email" validate:"required,email"`
}

func (ValidateOtp) TableName() string {
	return "User"
}
