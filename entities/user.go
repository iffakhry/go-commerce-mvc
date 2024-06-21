package entities

import "time"

type User struct {
	Id        uint
	Email     string
	Password  string
	Name      string
	Role      string
	CreatedAt time.Time
	UpdatedAt time.Time
}
