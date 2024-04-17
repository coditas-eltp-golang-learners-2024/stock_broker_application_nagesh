package repo

import (
	"Stock_broker_application/src/app/authentication/constants"
	"Stock_broker_application/src/app/authentication/models"
	"time"
)

func ValidateOtp(user models.ValidateOtp) error {
	var fetchedUser models.ValidateOtp
	if err := db.Where("email=?", user.Email).First(&fetchedUser).Error; err != nil {
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
