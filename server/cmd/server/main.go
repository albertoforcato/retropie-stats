package main

import (
	"os"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"github.com/albertoforcato/retropie-stats/db"
)

func main() {
	// Connect to the database
	dbConnection, err := gorm.Open(sqlite.Open(os.Getenv("SQLITE_DB_NAME")), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	// Execute the migrations
	err = db.ExecuteMigrations(dbConnection)
	if err != nil {
		panic(err)
	}
}
