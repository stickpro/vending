package handlers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/stickpro/vending/internal/domain"
	"net/http"
)

type OrderStoreRequest struct {
	ProductId  int
	TelegramId int
}

func (h *Handler) OrderStore(c *gin.Context) {
	//var order domain.Order
	var request OrderStoreRequest
	fmt.Println(request)
	if err := c.ShouldBindJSON(&request); err == nil {
		user, err := h.services.Users.FindByTgId(request.TelegramId)
		if err == nil {
			newResponse(c, http.StatusInternalServerError, err.Error())
			return
		}
		product, err := h.services.Products.GetById(request.ProductId)
		if err == nil {
			newResponse(c, http.StatusInternalServerError, err.Error())
			return
		}
		order := domain.Order{ProductId: int(product.ID), UserId: int(user.ID), ProductName: product.Name}
		newOrder, err := h.services.Orders.Create(order)

		if err != nil {
			newResponse(c, http.StatusInternalServerError, err.Error())
			return
		}
		c.JSON(http.StatusOK, newOrder)
	} else {
		newResponse(c, http.StatusUnprocessableEntity, err.Error())
	}

}

func (h *Handler) initOrdersRoutes(api *gin.RouterGroup) {
	orders := api.Group("/orders")
	{
		orders.POST("", h.OrderStore)
	}
}
