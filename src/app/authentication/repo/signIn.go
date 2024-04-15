package repo

import (
	"Stock_broker_application/src/app/authentication/models"
	utils "Stock_broker_application/src/app/authentication/utils/db"
)

func init() {
	var err error
	DB, err = utils.SetupDatabase()
	if err != nil {
		panic(err) // Handle the error appropriately
	}
}

func VerifySignInCredentials(user *models.UserSignIn) bool {
	query := "SELECT COUNT(*) FROM User WHERE email = ? and password  = ?"
	var data int64
	DB.QueryRow(query, user.Email, user.Password).Scan(&data)
	if data > 0 {
		return true
	}
	return false
}
