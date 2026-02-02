package products

import (
	"net/http"
	"strconv"

	"github.com/gabrielssssssssss/marketplace-telegram/internal/entity"
	"github.com/gabrielssssssssss/marketplace-telegram/internal/service"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

type ProductController struct {
	ProductService service.ProductService
}

func NewProductController(ProductService *service.ProductService) ProductController {
	return ProductController{ProductService: *ProductService}
}

func (controller ProductController) FetchProductByID(c *gin.Context) {
	productID, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		log.Error().
			Err(err).
			Str("component", "strconv.ParseInt").
			Str("product_id", c.Param("id")).
			Msg("Failed to fetch product request")

		c.JSON(http.StatusInternalServerError, gin.H{"error": "find_product_failed", "message": "InternalServerError"})
		return
	}

	findProduct := entity.Product{
		ID: &productID,
	}

	product, err := controller.ProductService.FindProductByID(&findProduct)
	if err != nil {
		log.Error().
			Err(err).
			Str("component", "controller.ProductService.FindProductByID").
			Int64("product_id", productID).
			Msg("Failed to fetch product request")

		c.JSON(http.StatusInternalServerError, gin.H{"error": "find_product_failed", "message": "InternalServerError"})
		return
	}

	log.Info().
		Str("status_code", "200").
		Int64("product_id", productID).
		Msg("Fetch product request processed successfully")

	c.JSON(http.StatusOK, gin.H{"message": "success", "data": product})
}
