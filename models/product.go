package models

import (
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	UserID      uint
	Name        string
	Price       float64
	Stock       uint
	Description string
	User        User
}
