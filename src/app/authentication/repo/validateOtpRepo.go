package repo

import (
	"Stock_broker_application/src/app/authentication/constants"
	"Stock_broker_application/src/app/authentication/models"
	"time"

	"gorm.io/gorm"
)

type ValidateOtpRepository interface {
	ValidateOtp(user models.ValidateOtp) error
}

func NewValidateOtpInstance(db *gorm.DB) *UserDBRepo {
	return &UserDBRepo{db: db}
}

func (repo *UserDBRepo) ValidateOtp(user models.ValidateOtp) error {
	var fetchedUser models.ValidateOtp
	if err := repo.db.Where("email=?", user.Email).First(&fetchedUser).Error; err != nil {
		return err
	}
	if fetchedUser.Otp != user.Otp {
		return constants.ErrOtpVerificationFailed
	}
	if fetchedUser.OtpExp.Before(time.Now()) {
		return constants.ErrOtpExpired
	}
	return nil
}
