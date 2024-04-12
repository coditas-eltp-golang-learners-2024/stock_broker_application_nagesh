package service

import (
	"Stock_broker_application/src/app/authentication/constants"
	"Stock_broker_application/src/app/authentication/models"
	"Stock_broker_application/src/app/authentication/repo"
)

func SignUp(user *models.User) error {

	result, err := repo.CheckUserExists(user)
	if err != nil {
		return constants.ErrInternalServer
	}

	if result > 0 {
		return constants.ErrUserExists
	}

	if err := repo.InsertUserIntoDB(user); err != nil {
		return constants.ErrInternalServer
	}

	return nil
}
