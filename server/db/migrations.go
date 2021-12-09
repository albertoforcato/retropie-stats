package db

import "gorm.io/gorm"

// ExecuteMigrations executes all migrations.
func ExecuteMigrations(db *gorm.DB) (err error) {
	err = db.AutoMigrate(&Game{}, &Entry{})
	if err != nil {
		return
	}

	return
}
