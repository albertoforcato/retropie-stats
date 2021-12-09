package main

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"github.com/albertoforcato/retropie-stats/db"
)

func main() {
	// Connect to the database
	dbConnection, err := gorm.Open(sqlite.Open("gorm.db"), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
	})

	// Execute the migrations
	err = db.ExecuteMigrations(dbConnection)
	if err != nil {
		panic(err)
	}
}
