package models

import (
	"gorm.io/gorm"
)

type Course struct {
	gorm.Model
	Title       string `json:"title"`
	Description string `json:"description"`
}

func Migrate(db *gorm.DB) {
	db.AutoMigrate(&Course{})
}
