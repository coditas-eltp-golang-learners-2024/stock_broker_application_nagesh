package repo

import (
	"Stock_broker_application/src/app/authentication/models"
	"time"
)

func GetUserByEmail(user *models.UserSignIn) (*models.UserSignIn, error) {
	var fetchedUser models.UserSignIn
	if err := db.Where("email=?", user.Email).First(&fetchedUser).Error; err != nil {
		return nil, err
	}
	return &fetchedUser, nil
}

func SetOtpAndOtpExp(user *models.UserSignIn, otp int, otpExp time.Time) error {
	err := db.Model(user).
		Where("email = ?", user.Email).
		Updates(map[string]interface{}{
			"otp":     otp,
			"otp_exp": otpExp,
		}).Error
	if err != nil {
		return err
	}
	return nil
}
