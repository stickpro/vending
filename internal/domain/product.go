package domain

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	Name   string `json:"name" gorm:"not null"`
	Url    string `json:"url" gorm:""`
	Status bool   `json:"status" gorm:"not null;default:true"`
	Code   uint8  `json:"code"`
	Image  string `json:"image"`
	Count  int    `json:"count"`
}
