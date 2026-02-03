package carts

import (
	"github.com/gabrielssssssssss/marketplace-telegram/internal/controller/api/v1/middlewares"
	"github.com/gin-gonic/gin"
)

func (controller *CartController) Route(rg *gin.RouterGroup) {
	rg.GET("/users/:id/carts", middlewares.Grant(), controller.FetchCartsByUserID)
	rg.POST("/carts", middlewares.Grant(), controller.InsertCart)
	rg.GET("/carts/:id", middlewares.Grant(), controller.FetchCartByID)
	rg.PUT("/carts/:id", middlewares.Grant(), controller.EditCartByID)
	rg.DELETE("/carts/:id", middlewares.Grant(), controller.DiscardCartByID)
}
