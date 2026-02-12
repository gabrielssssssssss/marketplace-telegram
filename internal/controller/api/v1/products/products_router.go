package products

import (
	"github.com/gabrielssssssssss/marketplace-telegram/internal/controller/api/v1/middlewares"
	"github.com/gin-gonic/gin"
)

func (controller *ProductController) Route(rg *gin.RouterGroup) {
	products := rg.Group("/products", middlewares.Authorization)
	{
		products.GET("", controller.GetAll)
		products.GET("/:id", controller.GetPublic)
		products.POST("", controller.Register, middlewares.Grant())
		products.GET("/:id/hidden", controller.GetPrivate, middlewares.Grant())
		products.PUT("/:id", controller.Modify, middlewares.Grant())
		products.DELETE("/:id", controller.Remove, middlewares.Grant())
	}
}
