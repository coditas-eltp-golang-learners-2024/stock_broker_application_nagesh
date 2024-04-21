package db

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func SetupGorm() (*gorm.DB, error) {
	dsn := "nagesh:SecureP@ssw0rd@tcp(localhost:3306)/stock_broker_application?parseTime=true"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return db, nil
}
