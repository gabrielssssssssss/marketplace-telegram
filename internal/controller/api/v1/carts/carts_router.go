package carts

import (
	"github.com/gabrielssssssssss/marketplace-telegram/internal/controller/api/v1/middlewares"
	"github.com/gin-gonic/gin"
)

func (controller *CartController) Route(rg *gin.RouterGroup) {
	carts := rg.Group("/carts", middlewares.Authorization, middlewares.Grant())
	{
		carts.POST("", controller.Register)
		carts.GET("", controller.GetCarts)
		carts.GET("/:id", controller.GetCart)
		carts.PUT("/:id", controller.Modify)
		carts.DELETE("/:id", controller.Remove)
	}

	rg.GET("/users/:id/carts",
		middlewares.Authorization,
		middlewares.Grant(),
		controller.GetUserCarts,
	)
}
