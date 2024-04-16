package repo

import (
	"Stock_broker_application/src/app/authentication/models"
	utils "Stock_broker_application/src/app/authentication/utils/db"

	"gorm.io/gorm"
)

var db *gorm.DB

func init() {
	var err error
	db, err = utils.SetupGorm()
	if err != nil {
		panic(err) // Handle the error appropriately
	}
}

func GetUserByEmailAndPassword(user *models.ChangePassword) (*models.User, error) {
	var fetchedUser models.User
	if err := db.Where("email=? AND password=?", user.Email, user.Password).First(&fetchedUser).Error; err != nil {
		return nil, err
	}
	return &fetchedUser, nil
}

func UpdateUserPassword(user *models.User, newUser *models.ChangePassword) error {
	if err := db.Model(user).Where("email=?", user.Email).Update("password", newUser.NewPassword).Error; err != nil {
		return err
	}
	return nil
}
