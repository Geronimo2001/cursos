package dtos

import "time"

type EnrollmentsDTO struct {
	UserID   uint      `json:"user_id"`
	CourseID uint      `json:"course_id"`
	Date     time.Time `json:"date"`
}
