package domain

import "gorm.io/gorm"

type Order struct {
	gorm.Model
	ProductId   int `json:"product_id" gorm:"not null"`
	UserId      int `json:"user_id" gorm:"not null"`
	User        User
	ProductName string `json:"product_name"`
}
