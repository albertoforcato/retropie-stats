package db

import (
	"time"

	"gorm.io/gorm"
)

type Game struct {
	gorm.Model
	ID              string
	CreatedAt       time.Time
	UpdatedAt       time.Time
	Path            string
	Name            string
	Year            string
	NumberOfPlayers uint
	Entries         []Entry
}

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
