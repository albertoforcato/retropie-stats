package v1

import (
	"gorm.io/gorm"

	"github.com/albertoforcato/retropie-stats/models"
)

type GameService struct{}

//GameList returns a list of games
func GameList(db *gorm.DB) (games []models.Game, err error) {
	db.Find(&games)
	return
}

//InsertNewGame inserts a new game into the database
func InsertNewGame(db *gorm.DB, game models.Game) (err error) {
	db.Create(&game)
	return
}
