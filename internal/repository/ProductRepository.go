package repository

import (
	"github.com/stickpro/vending/internal/domain"
	"github.com/stickpro/vending/pkg/logger"
	"gorm.io/gorm"
)

type productRepository struct {
	DB *gorm.DB
}

type Products interface {
	Save(domain.Product) (domain.Product, error)
	GetAll() ([]domain.Product, error)
	Migrate() error
}

func NewProductRepository(db *gorm.DB) Products {
	return productRepository{
		DB: db,
	}
}

func (p productRepository) Migrate() error {
	logger.Info("[ProductRepository]...Migrate")
	return p.DB.AutoMigrate(&domain.Product{})
}
func (p productRepository) Save(product domain.Product) (domain.Product, error) {
	logger.Info("[ProductRepository]...Save")
	err := p.DB.Create(&product).Error
	return product, err
}

func (p productRepository) GetAll() (products []domain.Product, err error) {
	logger.Info("[ProductRepository]...GetAll")
	err = p.DB.Find(&products).Error
	return products, err
}
