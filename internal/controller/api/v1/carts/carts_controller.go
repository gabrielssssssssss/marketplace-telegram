package carts

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gabrielssssssssss/marketplace-telegram/internal/entity"
	"github.com/gabrielssssssssss/marketplace-telegram/internal/model"
	"github.com/gabrielssssssssss/marketplace-telegram/internal/service"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

type CartController struct {
	CartService service.CartService
}

func NewCartController(CartService *service.CartService) CartController {
	return CartController{CartService: *CartService}
}

func (controller CartController) InsertCart(c *gin.Context) {
	var req model.Cart
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid_request", "message": "StatusBadRequest"})
		return
	}

	newCart := entity.Cart{
		UserID:    req.UserID,
		ProductID: req.ProductID,
	}

	cart, err := controller.CartService.RegisterCart(&newCart)
	if err != nil {
		log.Error().
			Err(err).
			Str("component", "controller.CartService.RegisterCart").
			Msg("Failed to insert cart request")

		c.JSON(http.StatusInternalServerError, gin.H{"error": "insert_cart_failed", "message": "InternalServerError"})
		return
	}

	log.Info().
		Str("status_code", "201").
		Msg("Insert cart request processed successfully")

	c.JSON(http.StatusCreated, gin.H{"message": "success", "data": cart})
}

func (controller CartController) FetchCartByID(c *gin.Context) {
	cartID := c.Param("id")
	findCart := entity.Cart{
		ID: cartID,
	}

	cart, err := controller.CartService.GetCartByID(&findCart)
	if err != nil {
		log.Error().
			Err(err).
			Str("component", "controller.CartService.FindCartByID").
			Str("cart_id", cartID).
			Msg("Failed to fetch cart request")

		c.JSON(http.StatusInternalServerError, gin.H{"error": "find_cart_failed", "message": "InternalServerError"})
		return
	}

	log.Info().
		Str("status_code", "200").
		Str("cart_id", cartID).
		Msg("Fetch cart request processed successfully")

	c.JSON(http.StatusOK, gin.H{"message": "success", "data": cart})
}

func (controller CartController) FetchCartsByUserID(c *gin.Context) {
	params := c.Param("id")

	userID, err := strconv.ParseInt(params, 10, 64)
	if err != nil {
		log.Error().
			Err(err).
			Str("component", "strconv.ParseInt").
			Int64("user_id", userID).
			Msg("Failed to fetch cart request")

		c.JSON(http.StatusInternalServerError, gin.H{"error": "find_cart_failed", "message": "InternalServerError"})
		return
	}

	findCarts := entity.Cart{
		UserID: userID,
	}

	carts, err := controller.CartService.GetCartsByUserID(&findCarts)
	if err != nil {
		log.Error().
			Err(err).
			Str("component", "controller.CartService.FindCartsByUserID").
			Int64("user_id", userID).
			Msg("Failed to fetch cart request")

		c.JSON(http.StatusInternalServerError, gin.H{"error": "find_cart_failed", "message": "InternalServerError"})
		return
	}

	log.Info().
		Str("status_code", "200").
		Int64("user_id", userID).
		Msg("Fetch cart request processed successfully")

	c.JSON(http.StatusOK, gin.H{"message": "success", "data": carts})
}

func (controller CartController) EditCartByID(c *gin.Context) {
	var req model.Cart
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid_request", "message": "StatusBadRequest"})
		return
	}

	updateCart := entity.Cart{
		ID:        c.Param("id"),
		UserID:    req.UserID,
		ProductID: req.ProductID,
		UpdatedAt: time.Now(),
	}

	_, err := controller.CartService.ModifyCartByID(&updateCart)
	if err != nil {
		log.Error().
			Err(err).
			Str("component", "controller.CartService.ModifyCartByID").
			Str("cart_id", c.Param("id")).
			Msg("Failed to edit cart request")

		c.JSON(http.StatusInternalServerError, gin.H{"error": "edit_cart_failed", "message": "InternalServerError"})
		return
	}

	log.Info().
		Str("status_code", "204").
		Str("cart_id", c.Param("id")).
		Msg("Edit cart request processed successfully")

	c.Status(http.StatusNoContent)
}

func (controller CartController) DiscardCartByID(c *gin.Context) {
	cartID := c.Param("id")
	removeCart := entity.Cart{
		ID: cartID,
	}

	_, err := controller.CartService.RemoveCartByID(&removeCart)
	if err != nil {
		log.Error().
			Err(err).
			Str("component", "controller.CartService.RemoveCartByID").
			Str("cart_id", cartID).
			Msg("Failed to fetch cart request")

		c.JSON(http.StatusInternalServerError, gin.H{"error": "discard_cart_failed", "message": "InternalServerError"})
		return
	}

	log.Info().
		Str("status_code", "204").
		Str("cart_id", cartID).
		Msg("Discard cart request processed successfully")

	c.Status(http.StatusNoContent)
}
