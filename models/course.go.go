package models

import (
	"time"
)

type User struct {
	ID        uint   `gorm:"primaryKey"`
	Name      string `gorm:"size:255"`
	Email     string `gorm:"size:255;unique"`
	Role      string `gorm:"size:255"`
	Password  string `gorm:"size:255"` // Atributo que no se transfiere
	createdat time.Time
	updatedat time.Time
}

type Course struct {
	ID          uint    `gorm:"primaryKey"`
	Name        string  `gorm:"size:255"`
	Length      float32 `gorm:"size:255"`
	Keywords    string  `gorm:"size:255"`
	Desc        string  `gorm:"size:255"`
	Req         string  `gorm:"size:255"` //requerimientos
	Teachername string  `gorm:"size:255"`
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
