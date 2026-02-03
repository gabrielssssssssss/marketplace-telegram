package orders

import (
	"github.com/gabrielssssssssss/marketplace-telegram/internal/controller/api/v1/middlewares"
	"github.com/gin-gonic/gin"
)

func (controller *OrderController) Route(rg *gin.RouterGroup) {
	rg.GET("/users/:id/carts", middlewares.Grant(), controller.FetchOrdersByUserID)
	rg.POST("/orders", middlewares.Grant(), controller.InsertOrder)
	rg.GET("/orders/:id", middlewares.Grant(), controller.FetchOrderByID)
	rg.DELETE("/orders/:id", middlewares.Grant(), controller.DiscardOrderByID)
}
