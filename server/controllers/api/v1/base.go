package v1

import (
	"gorm.io/gorm"
)

type PublicController struct {
	db *gorm.DB
}

func NewPublicController(db *gorm.DB) *PublicController {
	return &PublicController{db: db}
}
