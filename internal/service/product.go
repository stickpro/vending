package service

import (
	"github.com/stickpro/vending/internal/domain"
	"github.com/stickpro/vending/internal/repository"
)

type ProductsService struct {
	repository repository.Products
}

type ProductServiceInterface interface {
	LoadAll() ([]domain.Product, error)
	Create(product domain.Product) (domain.Product, error)
	GetById(id int) (domain.Product, error)
}

func NewProductsService(repository repository.Products) *ProductsService {
	return &ProductsService{
		repository: repository,
	}
}

func (p *ProductsService) LoadAll() ([]domain.Product, error) {
	return p.repository.GetAll()
}

func (p *ProductsService) Create(product domain.Product) (domain.Product, error) {
	return p.repository.Save(product)
}

func (p *ProductsService) GetById(id int) (domain.Product, error) {
	return p.repository.FindById(id)
}
