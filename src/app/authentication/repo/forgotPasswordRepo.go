package repo

import (
	"Stock_broker_application/src/app/authentication/models"

	"gorm.io/gorm"
)

type ForgotPasswordRepo interface {
	ValidateUserNamePanCard(user *models.ForgotPassword) error
	UpdatePassword(user *models.ForgotPassword) error
}

type ForgotPassword struct {
	db *gorm.DB
}

func NewForgetPasswordInstance(db *gorm.DB) *ForgotPassword {
	return &ForgotPassword{db: db}
}

func (repo *ForgotPassword) ValidateUserNamePanCard(user *models.ForgotPassword) error {

	if err := repo.db.Model(&models.ForgotPassword{}).Where("email = ? AND pan_card_number=?", user.Email, user.PanCardNumber).Where("pan_card_number = ?", user.PanCardNumber).First(user).Error; err != nil {
		return err
	}
	return nil
}
func (repo *ForgotPassword) UpdatePassword(user *models.ForgotPassword) error {

	if err := repo.db.Model(&models.ForgotPassword{}).Where("email = ? AND pan_card_number=?", user.Email, user.PanCardNumber).Update("password", user.NewPassword).Error; err != nil {
		return err
	}
	return nil
}
