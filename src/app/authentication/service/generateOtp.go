package service

import (
	"Stock_broker_application/src/app/authentication/models"
	"Stock_broker_application/src/app/authentication/repo"
	"Stock_broker_application/src/app/authentication/utils"
	"math/rand"
	"time"
)

func GenerateOtpService(user *models.UserSignIn) error {
	randSource := rand.NewSource(time.Now().UnixNano())
	random := rand.New(randSource)
	otp := random.Intn(10001)

	otpExpTime := CreateOtpExpTime(1)
	currentUser, err := repo.GetUserByEmail(user)
	utils.Logger.Println("Generated Otp", "otp", otp, "user_id", currentUser)
	repo.SetOtpAndOtpExp(currentUser, otp, otpExpTime)
	if err != nil {
		utils.Logger.Println("Error while generating Otp:", err)
		return err
	}
	return nil
}

func CreateOtpExpTime(hour int) time.Time {
	currentTimeUTC := time.Now().UTC()
	expTime := currentTimeUTC.Add(time.Hour * time.Duration(hour))
	utils.Logger.Println("Generated Otp Expiry Time:", expTime)
	return expTime
}
