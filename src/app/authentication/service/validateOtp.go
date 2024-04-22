package service

import (
	"Stock_broker_application/src/app/authentication/models"
	"Stock_broker_application/src/app/authentication/repo"
	"Stock_broker_application/src/app/authentication/utils"
)

type ValidateOtpService struct {
	ValidateOtpServiceRepository repo.ValidateOtpRepository
}

func NewValidateOtpService(validateOtpRepository repo.ValidateOtpRepository) *ValidateOtpService {
	return &ValidateOtpService{
		ValidateOtpServiceRepository: validateOtpRepository,
	}
}

func (service *ValidateOtpService) ValidateOtp(user models.ValidateOtp) error {
	utils.Logger.Println("INFO: ValidateOtpService function invoked")

	if err := service.ValidateOtpServiceRepository.ValidateOtp(user); err != nil {
		utils.Logger.Println("ERROR: Error in ValidateOtp function:", err)
		return err
	}
	utils.Logger.Println("INFO: Validation successful")
	return nil
}
