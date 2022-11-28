package repository

import (
	"github.com/stickpro/vending/internal/domain"
	"github.com/stickpro/vending/pkg/logger"
	"gorm.io/gorm"
)

type orderRepository struct {
	DB *gorm.DB
}

type Orders interface {
	Save(order domain.Order) (domain.Order, error)
	GetAll() ([]domain.Order, error)
	Migrate() error
}

func NewOrderRepository(db *gorm.DB) Orders {
	return orderRepository{
		DB: db,
	}
}

func (o orderRepository) Migrate() error {
	logger.Info("[OrderRepository]...Migrate")
	return o.DB.AutoMigrate(&domain.Order{})
}

func (o orderRepository) Save(order domain.Order) (domain.Order, error) {
	logger.Info("[OrderRepository]...Save")
	err := o.DB.Create(&order).Error
	return order, err
}

func (o orderRepository) GetAll() (orders []domain.Order, err error) {
	logger.Info("[OrderRepository]...GetAll")
	err = o.DB.Find(&orders).Error
	return orders, err
}
