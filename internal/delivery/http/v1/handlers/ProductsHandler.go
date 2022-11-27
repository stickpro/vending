package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/stickpro/vending/internal/domain"
	"net/http"
)

func (h *Handler) ProductsIndex(c *gin.Context) {
	products, err := h.services.Products.LoadAll()
	if err != nil {
		newResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, products)
}
func (h *Handler) ProductStore(c *gin.Context) {
	var product domain.Product
	if err := c.ShouldBindJSON(&product); err == nil {
		_, err := h.services.Products.Create(product)
		if err != nil {
			newResponse(c, http.StatusInternalServerError, err.Error())
			return
		}
		c.JSON(http.StatusOK, product)
	} else {
		newResponse(c, http.StatusUnprocessableEntity, err.Error())
	}
}

func (h *Handler) initProductsRoutes(api *gin.RouterGroup) {
	products := api.Group("/products")
	{
		products.GET("", h.ProductsIndex)
		products.POST("", h.ProductStore)
	}
}
