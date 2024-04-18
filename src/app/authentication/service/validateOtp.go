package service

import (
	"Stock_broker_application/src/app/authentication/models"
	"Stock_broker_application/src/app/authentication/repo"
	"Stock_broker_application/src/app/authentication/utils"
)

func ValidateOtpService(user models.ValidateOtp) error {
	utils.Logger.Println("INFO: ValidateOtpService function invoked")

	if err := repo.ValidateOtp(user); err != nil {
		utils.Logger.Println("ERROR: Error in ValidateOtp function:", err)
		return err
	}
	utils.Logger.Println("INFO: Validation successful")
	return nil
}
