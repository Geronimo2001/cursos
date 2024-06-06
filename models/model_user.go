package models

import (
	"strings"
)

type User struct {
	ID       uint   `gorm:"primaryKey;autoIncrement" json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Role     string `json:"role"`
	Password string `json:"password"`
}

func (u *User) FirstName() string {
	parts := strings.Split(u.Name, " ")
	if len(parts) > 0 {
		return parts[0]
	}
	return ""
}

func (u *User) LastName() string {
	parts := strings.Split(u.Name, " ")
	if len(parts) > 1 {
		return parts[len(parts)-1]
	}
	return ""
}
