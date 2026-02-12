package products

import (
	"github.com/gabrielssssssssss/marketplace-telegram/internal/controller/api/v1/middlewares"
	"github.com/gin-gonic/gin"
)

func (controller *ProductController) Route(rg *gin.RouterGroup) {
	products := rg.Group("/products", middlewares.Authorization)
	{
		products.GET("", controller.GetAll)
		products.POST("", middlewares.Grant(), controller.Register)
		products.GET("/:id", controller.GetPublic)
		products.PUT("/:id", middlewares.Grant(), controller.Modify)
		products.DELETE("/:id", middlewares.Grant(), controller.Remove)
		products.GET("/:id/hidden", middlewares.Grant(), controller.GetPrivate)
	}
}
