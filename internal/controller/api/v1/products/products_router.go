package products

import (
	"github.com/gabrielssssssssss/marketplace-telegram/internal/controller/api/v1/middlewares"
	"github.com/gin-gonic/gin"
)

func (controller *ProductController) Route(rg *gin.RouterGroup) {
	rg.POST("/products", middlewares.Grant(), controller.InsertProduct)
	rg.GET("/products/:id", middlewares.Grant(), controller.FetchProductByID)
	rg.PUT("/products/:id", middlewares.Grant(), controller.EditProductByID)
	// rg.DELETE("/users", middlewares.Grant(), controller.DiscardUserByID)
	// rg.PUT("/users", middlewares.Grant(), controller.EditUserByID)
}
