package service

import (
	"Stock_broker_application/src/app/authentication/models"
	"Stock_broker_application/src/app/authentication/repo"
	"Stock_broker_application/src/app/authentication/utils"
)

func ChangePasswordService(user *models.ChangePassword) (string, error) {

	utils.Logger.Println("Fetching user from database")
	fetchedUser, err := repo.GetUserByEmailAndPassword(user)
	if err != nil {
		utils.Logger.Println("Error fetching user from database", err)
		return "", err
	}
	utils.Logger.Println("Updating user password in database")
	result := repo.UpdateUserPassword(fetchedUser, user)
	if result != nil {
		utils.Logger.Println("Error updating user password in database", err)
		return "", result
	}
	utils.Logger.Println("Password Changed Successfully")
	return "Password Changed Successfully", nil
}
