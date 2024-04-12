package constants

import "errors"

var (
	ErrValidationFailed = errors.New("Validation failed")
	ErrUserExists       = errors.New("User already exists")
	ErrInternalServer   = errors.New("Internal server error")
)
