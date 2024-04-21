package utils

import (
	"Stock_broker_application/src/app/authentication/models"
	"strings"

	"github.com/go-playground/validator"
)

func SignUpValidations(user *models.User) error {
	validate := validator.New()

	err := validate.Struct(user)
	if err != nil {
		return err // Validation error occurred
	}

	return nil // No validation error
}
func ValidateCredentials(creds models.UserSignIn) error {
	validate := validator.New()

	validate.RegisterValidation("containsDigit", func(fl validator.FieldLevel) bool {
		password := fl.Field().String()
		for _, char := range password {
			if char >= '0' && char <= '9' {
				return true
			}
		}
		return false
	})

	validate.RegisterValidation("containsSpecial", func(fl validator.FieldLevel) bool {
		password := fl.Field().String()
		specialChars := "!@#$%^&*()-_=+[]{};:'\"|\\<>,./?"
		for _, char := range password {
			if strings.ContainsRune(specialChars, char) {
				return true
			}
		}
		return false
	})

	// Perform validation
	err := validate.Struct(creds)
	if err != nil {
		return err // Validation error occurred
	}

	return nil // No validation error
}
