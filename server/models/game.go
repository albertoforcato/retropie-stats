package models

import (
	"time"

	"gorm.io/gorm"
)

type Game struct {
	Path            string `gorm:"primaryKey"`
	CreatedAt       time.Time
	UpdatedAt       time.Time
	DeletedAt       gorm.DeletedAt
	Name            string
	Year            string
	NumberOfPlayers uint
	Entries         []Entry `gorm:"foreignKey:GamePath;references:Path"`
}
