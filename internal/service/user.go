package service

import "github.com/stickpro/vending/internal/repository"

type UsersService struct {
	repository repository.Users
}

type UserServiceInterface interface {
}

func NewUsersService(repository repository.Users) *UsersService {
	return &UsersService{
		repository: repository,
	}
}
