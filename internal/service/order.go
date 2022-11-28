package service

import (
	"github.com/stickpro/vending/internal/domain"
	"github.com/stickpro/vending/internal/repository"
)

type OrdersService struct {
	repository repository.Orders
}

type OrdersServiceInterface interface {
	LoadAll() ([]domain.Order, error)
	Create(order domain.Order) (domain.Order, error)
}

func NewOrderService(repository repository.Orders) *OrdersService {
	return &OrdersService{
		repository: repository,
	}
}

func (o *OrdersService) LoadAll() ([]domain.Order, error) {
	return o.repository.GetAll()
}

func (o *OrdersService) Create(order domain.Order) (domain.Order, error) {
	return o.repository.Save(order)
}
