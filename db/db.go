package db

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql" // Import MySQL driver
)

var db *sql.DB

// InitDB initializes the database connection
func InitDB(connectionString string) error {
	// Set up database connection parameters
	// Open a new database connection
	conn, err := sql.Open("mysql", connectionString)
	if err != nil {
		return err
	}

	// Ping the database to check if it's reachable
	err = conn.Ping()
	if err != nil {
		return err
	}

	fmt.Println("Connected to database!")
	db = conn
	return nil
}

// GetDB returns the database connection instance
func GetDB() *sql.DB {
	return db
}

func CloseDB() {
	db.Close()
}
