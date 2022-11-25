package repository

import (
	"gorm.io/gorm"
)

type Repositories struct {
	Users    Users
	Products Products
}

func NewRepositories(db *gorm.DB) *Repositories {
	return &Repositories{
		Users:    NewUserRepository(db),
		Products: NewProductRepository(db),
	}
}

func Migrate(r *Repositories) {
	_ = r.Users.Migrate()
	_ = r.Products.Migrate()
}
