package service

import (
	"Stock_broker_application/src/app/authentication/constants"
	"Stock_broker_application/src/app/authentication/models"
	"Stock_broker_application/src/app/authentication/repo"
)

func SignInService(user *models.UserSignIn) error {
	result := repo.VerifySignInCredentials(user)
	if result == false {
		return constants.ErrEmailOrPasswordVerificationFailed
	}
	return nil

}
