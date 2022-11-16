package repository

import (
	"gorm.io/gorm"
)

type Repositories struct {
	Users UserRepository
}

func NewRepositories(db *gorm.DB) *Repositories {
	return &Repositories{
		Users: NewUserRepository(db),
	}
}
