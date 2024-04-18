package service

import (
	"Stock_broker_application/src/app/authentication/constants"
	"Stock_broker_application/src/app/authentication/models"
	"Stock_broker_application/src/app/authentication/repo"
	"Stock_broker_application/src/app/authentication/utils"
)

func SignUpService(user *models.User) error {

	utils.Logger.Println("SignUpService function invoked")

	result, err := repo.CheckUserExists(user)
	if err != nil {
		utils.Logger.Println("Error in CheckUserExists function:", err)
		return constants.ErrInternalServer
	}

	if result > 0 {
		utils.Logger.Println("User already exists")
		return constants.ErrUserExists
	}

	if err := repo.InsertUserIntoDB(user); err != nil {
		utils.Logger.Println("Error in InsertUserIntoDB function:", err)
		return constants.ErrInternalServer
	}

	utils.Logger.Println("Successfully inserted user into DB")

	return nil
}
