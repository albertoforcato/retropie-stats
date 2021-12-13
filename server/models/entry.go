package models

import (
	"gorm.io/gorm"
)

const StartGameEvent = "start_game"
const EndGameEvent = "end_game"

type Entry struct {
	gorm.Model
	System      string
	Emulator    string
	EmulatorCmd string
	Event       string
	GamePath    string
}
