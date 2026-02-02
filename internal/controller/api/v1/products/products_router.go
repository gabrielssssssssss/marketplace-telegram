package products

import (
	"github.com/gabrielssssssssss/marketplace-telegram/internal/controller/api/v1/middlewares"
	"github.com/gin-gonic/gin"
)

func (controller *ProductController) Route(rg *gin.RouterGroup) {
	rg.GET("/products/:id", middlewares.Grant(), controller.FetchProductByID)
	// rg.DELETE("/users", middlewares.Grant(), controller.DiscardUserByID)
	// rg.PUT("/users", middlewares.Grant(), controller.EditUserByID)
}
