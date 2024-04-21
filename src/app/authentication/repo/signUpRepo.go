package repo

import (
	"Stock_broker_application/src/app/authentication/models"

	"gorm.io/gorm"
)

type UserSingUpRepository interface {
	CheckUserExists(user *models.User) (int, error)
	InsertUserIntoDB(user *models.User) error
}

func NewUserSignUpInstance(db *gorm.DB) *UserDBRepo {
	return &UserDBRepo{db: db}
}

func (repo *UserDBRepo) CheckUserExists(user *models.User) (int, error) {
	var count int64
	err := repo.db.Model(&models.User{}).Where("email = ? OR pan_card_number = ?", user.Email, user.PancardNumber).Count(&count).Error
	if err != nil {
		return 0, err
	}
	return int(count), nil
}
func (repo *UserDBRepo) InsertUserIntoDB(user *models.User) error {
	if err := repo.db.Create(user).Error; err != nil {
		return err
	}
	return nil
}
