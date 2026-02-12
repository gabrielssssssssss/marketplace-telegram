package products

import (
	"github.com/gabrielssssssssss/marketplace-telegram/internal/controller/api/v1/middlewares"
	"github.com/gin-gonic/gin"
)

func (controller *ProductController) Route(rg *gin.RouterGroup) {
	products := rg.Group("/products", middlewares.Authorization)
	{
		products.POST("", controller.InsertProduct, middlewares.Grant())
		products.GET("", controller.FetchAllProducts)
		products.GET("/:id", controller.FetchProductPublic)
		products.GET("/:id/hidden", controller.FetchProductPrivate, middlewares.Grant())
		products.PUT("/:id", controller.EditProductByID, middlewares.Grant())
		products.DELETE("/:id", controller.DiscardProductByID, middlewares.Grant())
	}
}
