package carts

import (
	"github.com/gabrielssssssssss/marketplace-telegram/internal/controller/api/v1/middlewares"
	"github.com/gin-gonic/gin"
)

func (controller *CartController) Route(rg *gin.RouterGroup) {
	carts := rg.Group("/carts", middlewares.Authorization, middlewares.Grant())
	{
		carts.POST("", controller.Create)
		carts.GET("", controller.Carts)
		carts.GET("/:id", controller.Cart)
		carts.PUT("/:id", controller.Update)
		carts.DELETE("/:id", controller.Delete)
	}

	rg.GET("/users/:id/carts",
		middlewares.Authorization,
		middlewares.Grant(),
		controller.UserCarts,
	)
}
