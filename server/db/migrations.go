package db

import (
	"gorm.io/gorm"

	"github.com/albertoforcato/retropie-stats/logger"
)

// ExecuteMigrations executes all migrations.
func ExecuteMigrations(db *gorm.DB) (err error) {
	logger.Log.Infof("Executing migrations...")
	err = db.AutoMigrate(&Game{}, &Entry{})
	if err != nil {
		return
	}
	logger.Log.Infof("Migrations executed successfully")

	return
}
