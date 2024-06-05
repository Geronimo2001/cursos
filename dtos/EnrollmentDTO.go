package dtos

import "time"

type Enrollment struct {
	UserID   uint      `json:"foreignKey:ID"`
	CourseID uint      `json:"foreignKey:ID"`
	Date     time.Time `json:"size:255"`
}
