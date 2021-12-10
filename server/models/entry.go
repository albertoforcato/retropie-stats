package models

import (
	"time"

	"gorm.io/gorm"
)

type Entry struct {
	gorm.Model
	ID          string
	CreatedAt   time.Time
	UpdatedAt   time.Time
	System      string
	Emulator    string
	EmulatorCmd string
	Event       string
	GameID      string
}
