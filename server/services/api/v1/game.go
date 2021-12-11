package v1

import "github.com/albertoforcato/retropie-stats/models"

type GameService struct {
	models.Game
}

func (g GameService) GameList() (games []models.Game, err error) {
	db := models.GetDB()

	db.Find(&games)
	return
}
