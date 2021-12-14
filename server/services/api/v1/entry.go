package v1

import (
	"gorm.io/gorm"

	"github.com/albertoforcato/retropie-stats/models"
)

//InsertStartGameEntry inserts a new start game entry
func InsertStartGameEntry(db *gorm.DB, entry models.Entry) (err error) {
	db.FirstOrCreate(&models.Game{Path: entry.GamePath})
	entry.Event = models.StartGameEvent
	db.Create(&entry)
	return
}

//InsertEndGameEntry inserts a new end game entry
func InsertEndGameEntry(db *gorm.DB, entry models.Entry) (err error) {
	db.FirstOrCreate(&models.Game{Path: entry.GamePath})
	entry.Event = models.EndGameEvent
	db.Create(&entry)
	return
}

//EntryList returns a list of entries
func EntryList(db *gorm.DB) (entries []models.Entry, err error) {
	db.Find(&entries)
	return
}
