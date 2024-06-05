package models

import "gorm.io/gorm"

type User struct {
	*gorm.Model
	Username string `json:"username"`
	PIN      string `json:"pin"`
	Email    string `json:"email"`
}
