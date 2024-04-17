package service

import (
	"Stock_broker_application/src/app/authentication/models"
	"Stock_broker_application/src/app/authentication/repo"
	"math/rand"
	"time"
)

func GenerateOtpService(user *models.UserSignIn) error {
	randSource := rand.NewSource(time.Now().UnixNano())
	random := rand.New(randSource)
	otp := random.Intn(10001)

	otpExpTime := CreateOtpExpTime(1)
	currentUser, err := repo.GetUserByEmail(user)
	repo.SetOtpAndOtpExp(currentUser, otp, otpExpTime)
	if err != nil {
		return err
	}
	return nil
}

func CreateOtpExpTime(hour int) time.Time {
	currentTimeUTC := time.Now().UTC()
	expTime := currentTimeUTC.Add(time.Hour * time.Duration(hour))
	return expTime
}
