package products

import (
	"github.com/gabrielssssssssss/marketplace-telegram/internal/controller/api/v1/middlewares"
	"github.com/gin-gonic/gin"
)

func (controller *ProductController) Route(rg *gin.RouterGroup) {
	products := rg.Group("/products", middlewares.Authorization, middlewares.Grant())
	{
		products.POST("", controller.InsertProduct)
		products.GET("", controller.FetchAllProducts)
		products.GET("/:id", controller.FetchProductByID)
		products.PUT("/:id", controller.EditProductByID)
		products.DELETE("/:id", controller.DiscardProductByID)
	}
}
