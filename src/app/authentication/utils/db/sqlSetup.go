//database stuct bind

package utils

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func SetupDatabase() (*sql.DB, error) {
	// dbConfig := utils.FetchDatabaseConfig()
	// dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbConfig.Database.Username, dbConfig.Database.Password, dbConfig.Database.Host, constants.DatabasePost, dbConfig.Database.DBName)

	//TODO: Fix it
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", "nagesh", "SecureP@ssw0rd", "localhost", "3306", "stock_broker_application")
	// Open a connection to the MySQL database
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, fmt.Errorf("error opening database connection: %v", err)
	}

	err = db.Ping()
	if err != nil {
		db.Close() // Close the connection before returning
		return nil, fmt.Errorf("error connecting to database: %v", err)
	}

	fmt.Println("Connected to MySQL database!")
	return db, err
}
