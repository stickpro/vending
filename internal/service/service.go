package service

import (
	"github.com/stickpro/vending/internal/repository"
)

type Services struct {
	Users    UserServiceInterface
	Products ProductServiceInterface
	Orders   OrdersServiceInterface
}
type Deps struct {
	Repository *repository.Repositories
}

func NewServices(deps Deps) *Services {
	userService := NewUsersService(deps.Repository.Users)
	productService := NewProductsService(deps.Repository.Products)
	orderService := NewOrderService(deps.Repository.Orders)
	return &Services{
		Users:    userService,
		Products: productService,
		Orders:   orderService,
	}
}
