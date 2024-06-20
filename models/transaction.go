package models

import "gorm.io/gorm"

type Transaction struct {
	gorm.Model
	UserID     uint
	ProductID  uint
	Price      int
	Quantity   int
	TotalPrice int
	User       User
	Product    Product
}
