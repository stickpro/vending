package service

import "github.com/stickpro/vending/internal/repository"

type ProductsService struct {
	repository repository.Products
}

type ProductServiceInterface interface {
}

func NewProductsService(repository repository.Products) *ProductsService {
	return &ProductsService{
		repository: repository,
	}
}
