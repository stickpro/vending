package router

import (
	"github.com/gorilla/mux"
	"github.com/stickpro/vending/internal/delivery/http/v1/handlers"
	"github.com/stickpro/vending/internal/service"
)

type Router struct {
	services *service.Services
}

func NewRouter(services *service.Services) *Router {
	return &Router{
		services: services,
	}
}

func (r *Router) Init() *mux.Router {
	route := mux.NewRouter()

	route.HandleFunc("/", handlers.HomePageIndex).Methods("GET")

	return route
}
