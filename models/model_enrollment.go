package models

import "time"

type Enrollment struct {
	UserID   uint      `json:"user_id"`
	CourseID uint      `json:"course_id"`
	Date     time.Time `json:"date"`
}
