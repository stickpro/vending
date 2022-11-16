package service

import (
	"github.com/stickpro/vending/internal/repository"
)

type Services struct {
	Users UserServiceInterface
}
type Deps struct {
	Repository *repository.Repositories
}

func NewServices(deps Deps) *Services {
	userService := NewUsersService(deps.Repository.Users)
	return &Services{
		Users: userService,
	}
}
