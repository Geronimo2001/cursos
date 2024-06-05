package models

import (
	"time"
)

type Course struct {
	ID          uint    `gorm:"primaryKey;autoIncrement" json:"id"`
	Name        string  `json:"name"`
	Length      float32 `json:"length"`
	Keywords    string  `json:"keywords"`
	Description string  `json:"description"`
	Req         string  `json:"req"`
	Teachername string  `json:"teachername"`
	createdat   time.Time
	updatedat   time.Time
}

type Enrollment struct {
	ID        uint   `gorm:"primaryKey"`
	UserID    uint   `gorm:"foreignKey:ID"`
	CourseID  uint   `gorm:"foreignKey:ID"`
	Date      string `gorm:"size:255"`
	createdat time.Time
	updatedat time.Time
}
