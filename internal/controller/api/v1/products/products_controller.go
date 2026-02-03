package products

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gabrielssssssssss/marketplace-telegram/internal/entity"
	"github.com/gabrielssssssssss/marketplace-telegram/internal/model"
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

func (controller ProductController) InsertProduct(c *gin.Context) {
	var req model.Product
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid_request", "message": "StatusBadRequest"})
		return
	}

	fmt.Println(req.Price)
	newProduct := entity.Product{
		Details: &req.Details,
		Price:   &req.Price,
	}

	product, err := controller.ProductService.RegisterProduct(&newProduct)
	if err != nil {
		log.Error().
			Err(err).
			Str("component", "controller.ProductService.RegisterProduct").
			Msg("Failed to insert product request")

		c.JSON(http.StatusInternalServerError, gin.H{"error": "insert_product_failed", "message": "InternalServerError"})
		return
	}

	log.Info().
		Str("status_code", "201").
		Msg("Insert product request processed successfully")

	c.JSON(http.StatusCreated, gin.H{"message": "success", "data": product})

}

func (controller ProductController) FetchProductByID(c *gin.Context) {
	productID := c.Param("id")
	findProduct := entity.Product{
		ID: productID,
	}

	product, err := controller.ProductService.FindProductByID(&findProduct)
	if err != nil {
		log.Error().
			Err(err).
			Str("component", "controller.ProductService.FindProductByID").
			Str("product_id", productID).
			Msg("Failed to fetch product request")

		c.JSON(http.StatusInternalServerError, gin.H{"error": "find_product_failed", "message": "InternalServerError"})
		return
	}

	log.Info().
		Str("status_code", "200").
		Str("product_id", productID).
		Msg("Fetch product request processed successfully")

	c.JSON(http.StatusOK, gin.H{"message": "success", "data": product})
}

func (controller ProductController) EditProductByID(c *gin.Context) {
	var req model.Product
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid_request", "message": "StatusBadRequest"})
		return
	}

	updateProduct := entity.Product{
		ID:        c.Param("id"),
		Details:   &req.Details,
		Price:     &req.Price,
		UpdatedAt: time.Now(),
	}

	_, err := controller.ProductService.ModifyProductByID(&updateProduct)
	if err != nil {
		log.Error().
			Err(err).
			Str("component", "controller.ProductService.ModifyProductByID").
			Str("product_id", c.Param("id")).
			Msg("Failed to edit product request")

		c.JSON(http.StatusInternalServerError, gin.H{"error": "edit_product_failed", "message": "InternalServerError"})
		return
	}

	log.Info().
		Str("status_code", "204").
		Str("product_id", c.Param("id")).
		Msg("Edit product request processed successfully")

	c.Status(http.StatusNoContent)
}
