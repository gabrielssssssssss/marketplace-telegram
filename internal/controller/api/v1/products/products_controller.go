package products

import (
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

// InsertProduct   godoc
// @Summary        Insert product data
// @Description    post product data
// @Tags           products
// @Accept         json
// @Produce        json
// @Param          Authorization  header  string  true  "Insert your admin JWT token"
// @Success        201
// @Failure        400  {object}  model.Error
// @Failure        500  {object}  model.Error
// @Router         /products/ [post]
func (controller ProductController) InsertProduct(c *gin.Context) {
	var req model.Product
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, model.Error{Error: "invalid_request", Message: "StatusBadRequest"})
		return
	}

	product, err := controller.ProductService.RegisterProduct(&entity.Product{
		Details: &req.Details,
		Price:   &req.Price,
	})
	if err != nil {
		log.Error().
			Err(err).
			Str("component", "controller.ProductService.RegisterProduct").
			Msg("Failed to insert product request")

		c.JSON(http.StatusInternalServerError, model.Error{Error: "insert_product_failed", Message: "InternalServerError"})
		return
	}

	log.Info().
		Str("status_code", "201").
		Msg("Insert product request processed successfully")

	c.JSON(http.StatusCreated, gin.H{"message": "success", "data": product})
}

// FetchProductByID godoc
// @Summary         Get product data
// @Description     get product data
// @Tags            products
// @Accept          json
// @Produce         json
// @Param           Authorization  header  string  true  "Insert your admin JWT token"
// @Success         200  {object}  model.ProductResponse
// @Failure         400  {object}  model.Error
// @Failure         500  {object}  model.Error
// @Router          /products/:id [get]
func (controller ProductController) FetchProductByID(c *gin.Context) {
	productID := c.Param("id")

	product, err := controller.ProductService.GetProductByID(&entity.Product{
		ID: productID,
	})
	if err != nil {
		log.Error().
			Err(err).
			Str("component", "controller.ProductService.FindProductByID").
			Str("product_id", productID).
			Msg("Failed to fetch product request")

		c.JSON(http.StatusInternalServerError, model.Error{Error: "find_product_failed", Message: "InternalServerError"})
		return
	}

	log.Info().
		Str("status_code", "200").
		Str("product_id", productID).
		Msg("Fetch product request processed successfully")

	c.JSON(http.StatusOK, model.ProductResponse{Message: "success", Data: *product})
}

// EditProductByID  godoc
// @Summary         Edit product data
// @Description     put product data
// @Tags            products
// @Accept          json
// @Produce         json
// @Param           Authorization  header  string  true  "Insert your admin JWT token"
// @Success         204
// @Failure         400  {object}  model.Error
// @Failure         500  {object}  model.Error
// @Router          /products/:id [put]
func (controller ProductController) EditProductByID(c *gin.Context) {
	productID := c.Param("id")

	var req model.Product
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, model.Error{Error: "invalid_request", Message: "StatusBadRequest"})
		return
	}

	_, err := controller.ProductService.ModifyProductByID(&entity.Product{
		ID:        productID,
		Details:   &req.Details,
		Price:     &req.Price,
		UpdatedAt: time.Now(),
	})
	if err != nil {
		log.Error().
			Err(err).
			Str("component", "controller.ProductService.ModifyProductByID").
			Str("product_id", productID).
			Msg("Failed to edit product request")

		c.JSON(http.StatusInternalServerError, model.Error{Error: "edit_product_failed", Message: "InternalServerError"})
		return
	}

	log.Info().
		Str("status_code", "204").
		Str("product_id", productID).
		Msg("Edit product request processed successfully")

	c.Status(http.StatusNoContent)
}

// DiscardProductByID godoc
// @Summary           Discard product data
// @Description       delete product data
// @Tags              products
// @Accept            json
// @Produce           json
// @Param             Authorization  header  string  true  "Insert your admin JWT token"
// @Success           204
// @Failure           400  {object}  model.Error
// @Failure           500  {object}  model.Error
// @Router            /products/:id [delete]
func (controller ProductController) DiscardProductByID(c *gin.Context) {
	productID := c.Param("id")

	_, err := controller.ProductService.RemoveProductByID(&entity.Product{
		ID: productID,
	})
	if err != nil {
		log.Error().
			Err(err).
			Str("component", "controller.ProductService.RemoveProductByID").
			Str("product_id", productID).
			Msg("Failed to fetch product request")

		c.JSON(http.StatusInternalServerError, gin.H{"error": "discard_product_failed", "message": "InternalServerError"})
		return
	}

	log.Info().
		Str("status_code", "204").
		Str("product_id", productID).
		Msg("Discard product request processed successfully")

	c.Status(http.StatusNoContent)
}
