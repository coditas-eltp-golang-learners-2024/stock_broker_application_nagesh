package service

import (
	"Stock_broker_application/src/app/authentication/models"
	"Stock_broker_application/src/app/authentication/repo"
)

func ChangePasswordService(user *models.ChangePassword) (string, error) {

	fetchedUser, err := repo.GetUserByEmailAndPassword(user)
	if err != nil {
		return "", err
	}
	result := repo.UpdateUserPassword(fetchedUser, user)
	if result != nil {
		return "", result
	}
	return "Password Changed Successfully", nil
}
