package entity

import "time"

type Product struct {
	Id          uint
	UserID      uint
	Name        string
	Price       float64
	Stock       uint
	Description string
	User        User
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
