package carts

import (
	"github.com/gabrielssssssssss/marketplace-telegram/internal/controller/api/v1/middlewares"
	"github.com/gin-gonic/gin"
)

func (controller *CartController) Route(rg *gin.RouterGroup) {
	carts := rg.Group("/carts", middlewares.Authorization, middlewares.Grant())
	{
		carts.POST("", controller.InsertCart)
		carts.GET("/:id", controller.FetchCartByID)
		carts.PUT("/:id", controller.EditCartByID)
		carts.DELETE("/:id", controller.DiscardCartByID)
	}

	rg.GET("/users/:id/carts",
		middlewares.Authorization,
		middlewares.Grant(),
		controller.FetchCartsByUserID,
	)
}
