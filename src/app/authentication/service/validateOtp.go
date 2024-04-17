package service

import (
	"Stock_broker_application/src/app/authentication/models"
	"Stock_broker_application/src/app/authentication/repo"
)

func ValidateOtpService(user models.ValidateOtp) error {
	if err := repo.ValidateOtp(user); err != nil {
		return err
	}
	return nil
}
