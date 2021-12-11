package models

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"github.com/albertoforcato/retropie-stats/logger"
)

var db *gorm.DB

const dbName = "retropie-stats.db"

func init() {
	// Connect to the database
	conn, err := gorm.Open(sqlite.Open(dbName), &gorm.Config{})
	if err != nil {
		logger.Log.Errorf("Error connecting to database: %v", err)
	}
	db = conn
	// Execute the migrations
	err = db.Debug().AutoMigrate(&Game{}, &Entry{})
	if err != nil {
		logger.Log.Errorf("Error executing migrations: %v", err)
	}
}

//GetDB function return the instance of db
func GetDB() *gorm.DB {
	return db
}
