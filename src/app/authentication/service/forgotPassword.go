package service

import (
	"Stock_broker_application/src/app/authentication/models"
	"Stock_broker_application/src/app/authentication/repo"
)

type ForgotPasswordService struct {
	ForgotPasswordRepo repo.ForgotPasswordRepo
}

func NewForgetPasswordService(forgotPasswordRepo repo.ForgotPasswordRepo) *ForgotPasswordService {
	return &ForgotPasswordService{
		ForgotPasswordRepo: forgotPasswordRepo,
	}
}

func (service *ForgotPasswordService) ForgotPassword(user *models.ForgotPassword) error {

	if err := service.ForgotPasswordRepo.ValidateUserNamePanCard(user); err != nil {
		return err
	}
	if err := service.ForgotPasswordRepo.UpdatePassword(user); err != nil {
		return err
	}
	return nil
}
