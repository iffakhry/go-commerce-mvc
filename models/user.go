package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	// ID        string `gorm:"primaryKey"`
	// CreatedAt time.Time
	// UpdatedAt time.Time
	// DeletedAt gorm.DeletedAt `gorm:"index"`
	Email    string `gorm:"unique"`
	Password string
	Name     string
	Role     string
}
