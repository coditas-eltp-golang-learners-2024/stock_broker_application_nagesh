package repo

import (
	"Stock_broker_application/src/app/authentication/models"
	utils "Stock_broker_application/src/app/authentication/utils/db"

	"gorm.io/gorm"
)

var database *gorm.DB

func init() {
	var err error
	database, err = utils.SetupGorm()
	if err != nil {
		panic(err) // Handle the error appropriately
	}
}

type ChangePasswordRepository interface {
	UpdateUserPassword(user *models.User, newUser *models.ChangePassword) error
	GetUserByEmailAndPassword(user *models.ChangePassword) (*models.User, error)
}

func NewChangePasswordInstance(db *gorm.DB) *UserDBRepo {
	return &UserDBRepo{db: db}
}

func (repo *UserDBRepo) GetUserByEmailAndPassword(user *models.ChangePassword) (*models.User, error) {
	var fetchedUser models.User
	if err := repo.db.Where("email=? AND password=?", user.Email, user.Password).First(&fetchedUser).Error; err != nil {
		return nil, err
	}
	return &fetchedUser, nil
}

func (repo *UserDBRepo) UpdateUserPassword(user *models.User, newUser *models.ChangePassword) error {
	if err := repo.db.Model(user).Where("email=?", user.Email).Update("password", newUser.NewPassword).Error; err != nil {
		return err
	}
	return nil
}
