package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) UsersPageIndex(c *gin.Context) {
	users, err := h.services.Users.LoadAll()
	if err != nil {
		newResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, users)
}

func (h *Handler) initUsersRoutes(api *gin.RouterGroup) {
	users := api.Group("/users")
	{
		users.GET("", h.UsersPageIndex)
	}
}
