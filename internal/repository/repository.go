package repository

import (
	"gorm.io/gorm"
)

type Repositories struct {
	Users    Users
	Products Products
	Orders   Orders
}

func NewRepositories(db *gorm.DB) *Repositories {
	return &Repositories{
		Users:    NewUserRepository(db),
		Products: NewProductRepository(db),
		Orders:   NewOrderRepository(db),
	}
}

func Migrate(r *Repositories) {
	_ = r.Users.Migrate()
	_ = r.Products.Migrate()
	_ = r.Orders.Migrate()
}
