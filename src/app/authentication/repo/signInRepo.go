// package repo

// import (
// 	"Stock_broker_application/src/app/authentication/models"
// )

// func VerifySignInCredentials(user *models.UserSignIn) bool {
// 	var count int64
// 	database.Model(&models.User{}).Where("email = ? AND password = ?", user.Email, user.Password).Count(&count)
// 	return count > 0
// }

package repo

import (
	"Stock_broker_application/src/app/authentication/constants"
	"Stock_broker_application/src/app/authentication/models"
	"log"
	"time"

	"gorm.io/gorm"
)

// UserRepository defines the interface for interacting with User data.
type UserSignInRepository interface {
	IsEmailExists(email string) bool
	IsPhoneNumberExists(phoneNumber string) bool
	IsPancardNumberExists(pancardNumber string) bool
	InsertCustomer(customer *models.User) error
	GetUserByEmail(email string) (*models.UserSignIn, error)
	SetOtpAndOtpExpiry(user *models.UserSignIn, otp int, otpExpiry time.Time) error
}

// UserDBRepository is an implementation of UserRepository using GORM for database interactions.
type UserDBRepo struct {
	db *gorm.DB
}

// NewUserRepository creates a new instance of UserDBRepository.
func NewUserInstance(db *gorm.DB) *UserDBRepo {
	return &UserDBRepo{db: db}
}

func (repo *UserDBRepo) IsEmailExists(email string) bool {
	var count int64
	if err := repo.db.Model(&models.User{}).Where("email = ?", email).Count(&count).Error; err != nil {
		return false
	}
	return count > 0
}

func (repo *UserDBRepo) IsPhoneNumberExists(phoneNumber string) bool {
	var count int64
	if err := repo.db.Model(&models.User{}).Where("phone_number = ?", phoneNumber).Count(&count).Error; err != nil {
		return false
	}
	return count > 0
}

func (repo *UserDBRepo) IsPancardNumberExists(pancardNumber string) bool {
	var count int64
	if err := repo.db.Model(&models.User{}).Where("pancard_number = ?", pancardNumber).Count(&count).Error; err != nil {
		return false
	}
	return count > 0
}

func (repo *UserDBRepo) InsertCustomer(customer *models.User) error {
	if err := repo.db.Create(customer).Error; err != nil {
		return err
	}
	return nil
}

func (repo *UserDBRepo) GetUserByEmail(email string) (*models.UserSignIn, error) {
	var customer models.UserSignIn
	if err := repo.db.Where("email = ?", email).First(&customer).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, constants.ErrRecordNotFound
		}
		return nil, err
	}
	return &customer, nil
}

func (repo *UserDBRepo) SetOtpAndOtpExpiry(user *models.UserSignIn, otp int, otpExp time.Time) error {
	log.Println("SetOtpAndOtpExpiry method called with user:", user, "otp:", otp, "otpExp:", otpExp)
	err := repo.db.Model(user).
		Where("email = ?", user.Email).
		Updates(map[string]interface{}{
			"otp":     otp,
			"otp_exp": otpExp,
		}).Error
	if err != nil {
		log.Println("Error updating user with otp and otpExp:", err)
		return err
	}
	log.Println("Successfully updated user with otp and otpExp")
	return nil
}
