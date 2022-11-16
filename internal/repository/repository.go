package repository

import (
	"gorm.io/gorm"
)

type Repositories struct {
	Users Users
}

func NewRepositories(db *gorm.DB) *Repositories {
	return &Repositories{
		Users: NewUserRepository(db),
	}
}
