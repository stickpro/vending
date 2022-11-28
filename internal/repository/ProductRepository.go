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
	FindByColumn(any, string) (domain.Product, error)
	FindById(int) (domain.Product, error)
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

func (p productRepository) FindByColumn(value any, columnName string) (domain.Product, error) {
	logger.Info("[UserRepository]... Find by column")
	var product domain.Product
	err := p.DB.Find(&product, columnName+" = ?", value).Error
	return product, err
}

func (p productRepository) FindById(value int) (domain.Product, error) {
	logger.Info("[UserRepository]... Find by id")
	var product domain.Product
	err := p.DB.First(&product, value).Error
	return product, err
}
