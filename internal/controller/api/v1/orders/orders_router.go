package orders

import (
	"github.com/gabrielssssssssss/marketplace-telegram/internal/controller/api/v1/middlewares"
	"github.com/gin-gonic/gin"
)

func (controller *OrderController) Route(rg *gin.RouterGroup) {
	orders := rg.Group("/orders", middlewares.Authorization, middlewares.Grant())
	{
		orders.POST("", controller.InsertOrder)
		orders.GET("/:id", controller.FetchOrderByID)
		orders.DELETE("/:id", controller.DiscardOrderByID)
	}
	rg.GET("/users/:id/orders",
		middlewares.Authorization,
		middlewares.Grant(),
		controller.FetchOrdersByUserID,
	)
}
