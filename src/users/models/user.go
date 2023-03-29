package models

import (
	"gorm.io/gorm"
)


type User struct {
	gorm.Model
	Name     string  `json:"name"`
	Email    string  `json:"email"`
	Username string  `json:"username"`
	Password string  `json:"password"`
	Address  string  `json:"address"`
	Phone    string  `json:"phone"`
	Role     string  `json:"role"`
}