package orders

import (
	"github.com/gabrielssssssssss/marketplace-telegram/internal/controller/api/v1/middlewares"
	"github.com/gin-gonic/gin"
)

func (controller *OrderController) Route(rg *gin.RouterGroup) {
	orders := rg.Group("/orders", middlewares.Authorization, middlewares.Grant())
	{
		orders.POST("", controller.Register)
		orders.GET("/:id", controller.GetOrder)
		orders.DELETE("/:id", controller.Remove)
	}
	rg.GET("/users/:id/orders",
		middlewares.Authorization,
		middlewares.Grant(),
		controller.GetUserOrders,
	)
}
