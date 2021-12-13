package models

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"github.com/albertoforcato/retropie-stats/logger"
)

const dbName = "retropie-stats.db"

//NewDbInstance function return the instance of db
func NewDbInstance() *gorm.DB {
	// Connect to the database
	conn, err := gorm.Open(sqlite.Open(dbName), &gorm.Config{})
	if err != nil {
		logger.Log.Errorf("Error connecting to database: %v", err)
	}
	// Execute the migrations
	err = conn.Debug().AutoMigrate(&Game{}, &Entry{})
	if err != nil {
		logger.Log.Errorf("Error executing migrations: %v", err)
	}
	return conn
}
