package constants

import "errors"

var (
	ErrValidationFailed                  = errors.New("Validation failed")
	ErrUserExists                        = errors.New("User already exists")
	ErrInternalServer                    = errors.New("Internal server error")
	ErrEmailOrPasswordVerificationFailed = errors.New("Email or password wrong")
	ErrOtpVerificationFailed             = errors.New("Otp verification failed")
	ErrOtpExpired                        = errors.New("Otp expired")
	ErrRecordNotFound                    = errors.New("Record not found")
	ErrUserNotExists                     = errors.New("User not exists")
	ErrInvalidCredentials                = errors.New("Invalid credentials")
)
