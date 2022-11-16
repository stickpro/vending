package domain

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Email     string `json:"email" gorm:"unique;not null"`
	Password  string `json:"password" gorm:"not null"`
	FirstName string `json:"first_name" gorm:"varchar(255);not null"`
	LastName  string `json:"last_name" gorm:"varchar(255);not null"`
	Status    bool   `json:"status" gorm:"not null;default:false"`
}
