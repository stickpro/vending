package repository

import (
	"github.com/stickpro/vending/internal/domain"
	"github.com/stickpro/vending/pkg/logger"
	"gorm.io/gorm"
)

type userRepository struct {
	DB *gorm.DB
}

type UserRepository interface {
	Save(domain.User) (domain.User, error)
	GetAll() ([]domain.User, error)
	Migrate() error
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return userRepository{
		DB: db,
	}
}

func (u userRepository) Migrate() error {
	logger.Info("[UserRepository]...Migrate")
	return u.DB.AutoMigrate(&domain.User{})
}

func (u userRepository) Save(user domain.User) (domain.User, error) {
	logger.Info("[UserRepository]...Save")
	err := u.DB.Create(&user).Error
	return user, err
}

func (u userRepository) GetAll() (users []domain.User, err error) {
	logger.Info("[UserRepository]...Get All")
	err = u.DB.Find(&users).Error
	return users, err
}
