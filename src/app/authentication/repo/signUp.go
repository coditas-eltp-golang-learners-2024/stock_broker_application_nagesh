package repo

import (
	"Stock_broker_application/src/app/authentication/models"
	utils "Stock_broker_application/src/app/authentication/utils/db"
	"database/sql"
)

var DB *sql.DB

func init() {
	var err error
	DB, err = utils.SetupDatabase()
	if err != nil {
		panic(err) // Handle the error appropriately
	}
}

func CheckUserExists(user *models.User) (int, error) {
	query := `
        SELECT COUNT(*) AS user_count
        FROM User
        WHERE email = ? OR pan_card_number = ?
    `
	var count int
	err := DB.QueryRow(query, user.Email, user.PancardNumber).Scan(&count)
	if err != nil {
		return 0, err
	}
	return count, nil
}
func InsertUserIntoDB(user *models.User) error {
	_, err := DB.Exec("INSERT INTO User (name, email, phone_number, pan_card_number, password) VALUES (?, ?, ?, ?, ?)",
		user.Name, user.Email, user.PhoneNumber, user.PancardNumber, user.Password)
	if err != nil {
		return err
	}
	return nil
}
