package models

import "github.com/jinzhu/gorm"

type Teachers struct {
	gorm.Model
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}
