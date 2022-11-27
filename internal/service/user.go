package service

import (
	"github.com/stickpro/vending/internal/domain"
	"github.com/stickpro/vending/internal/repository"
)

type UsersService struct {
	repository repository.Users
}

type UserServiceInterface interface {
	LoadAll() ([]domain.User, error)
}

func NewUsersService(repository repository.Users) *UsersService {
	return &UsersService{
		repository: repository,
	}
}

func (u *UsersService) LoadAll() ([]domain.User, error) {
	return u.repository.GetAll()
}
