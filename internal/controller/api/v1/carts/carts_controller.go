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

// InsertCart    godoc
// @Summary      Insert cart data
// @Description  post cart data
// @Tags         carts
// @Accept       json
// @Produce      json
// @Param        Authorization  header  string  true  "Insert your admin JWT token"
// @Success      201  {object}  model.CartResponse
// @Failure      400  {object}  model.Error
// @Failure      500  {object}  model.Error
// @Router         /carts/ [post]
func (controller CartController) InsertCart(c *gin.Context) {
	var req model.Cart
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, model.Error{Error: "invalid_request", Message: "StatusBadRequest"})
		return
	}

	cart, err := controller.CartService.RegisterCart(&entity.Cart{
		UserID:    req.UserID,
		ProductID: req.ProductID,
	})
	if err != nil {
		log.Error().
			Err(err).
			Str("component", "controller.CartService.RegisterCart").
			Msg("Failed to insert cart request")

		c.JSON(http.StatusInternalServerError, model.Error{Error: "insert_cart_failed", Message: "InternalServerError"})
		return
	}

	log.Info().
		Str("status_code", "201").
		Msg("Insert cart request processed successfully")

	c.JSON(http.StatusCreated, model.CartResponse{Message: "success", Data: *cart})
}

// FetchCartByID    godoc
// @Summary         Get cart data
// @Description     get cart data
// @Tags            carts
// @Accept          json
// @Produce         json
// @Param           Authorization  header  string  true  "Insert your admin JWT token"
// @Success         200  {object}  model.CartResponse
// @Failure         500  {object}  model.Error
// @Router          /carts/:id [get]
func (controller CartController) FetchCartByID(c *gin.Context) {
	cartID := c.Param("id")

	cart, err := controller.CartService.GetCartByID(&entity.Cart{
		ID: cartID,
	})
	if err != nil {
		log.Error().
			Err(err).
			Str("component", "controller.CartService.FindCartByID").
			Str("cart_id", cartID).
			Msg("Failed to fetch cart request")

		c.JSON(http.StatusInternalServerError, model.Error{Error: "find_cart_failed", Message: "InternalServerError"})
		return
	}

	log.Info().
		Str("status_code", "200").
		Str("cart_id", cartID).
		Msg("Fetch cart request processed successfully")

	c.JSON(http.StatusOK, model.CartResponse{Message: "success", Data: *cart})
}

// FetchCartsByUserID  godoc
// @Summary            Get user carts data
// @Description        get all carts from userID data
// @Tags               carts
// @Accept             json
// @Produce            json
// @Param              Authorization  header  string  true  "Insert your admin JWT token"
// @Success            200  {object}  model.CartsResponse
// @Failure            500  {object}  model.Error
// @Router             /users/:id/carts [get]
func (controller CartController) FetchCartsByUserID(c *gin.Context) {
	params := c.Param("id")

	userID, err := strconv.ParseInt(params, 10, 64)
	if err != nil {
		log.Error().
			Err(err).
			Str("component", "strconv.ParseInt").
			Int64("user_id", userID).
			Msg("Failed to fetch cart request")

		c.JSON(http.StatusInternalServerError, model.Error{Error: "find_cart_failed", Message: "InternalServerError"})
		return
	}

	carts, err := controller.CartService.GetCartsByUserID(&entity.Cart{
		UserID: userID,
	})
	if err != nil {
		log.Error().
			Err(err).
			Str("component", "controller.CartService.FindCartsByUserID").
			Int64("user_id", userID).
			Msg("Failed to fetch cart request")

		c.JSON(http.StatusInternalServerError, model.Error{Error: "find_cart_failed", Message: "InternalServerError"})
		return
	}

	log.Info().
		Str("status_code", "200").
		Int64("user_id", userID).
		Msg("Fetch cart request processed successfully")

	c.JSON(http.StatusOK, model.CartsResponse{Message: "success", Data: *carts})
}

// EditCartByID     godoc
// @Summary         Edit cart data
// @Description     put cart data
// @Tags            carts
// @Accept          json
// @Produce         json
// @Param           Authorization  header  string  true  "Insert your admin JWT token"
// @Success         204
// @Failure         400  {object}  model.Error
// @Failure         500  {object}  model.Error
// @Router          /carts/:id [put]
func (controller CartController) EditCartByID(c *gin.Context) {
	var req model.Cart
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, model.Error{Error: "invalid_request", Message: "StatusBadRequest"})
		return
	}

	_, err := controller.CartService.ModifyCartByID(&entity.Cart{
		ID:        c.Param("id"),
		UserID:    req.UserID,
		ProductID: req.ProductID,
		UpdatedAt: time.Now(),
	})
	if err != nil {
		log.Error().
			Err(err).
			Str("component", "controller.CartService.ModifyCartByID").
			Str("cart_id", c.Param("id")).
			Msg("Failed to edit cart request")

		c.JSON(http.StatusInternalServerError, model.Error{Error: "edit_cart_failed", Message: "InternalServerError"})
		return
	}

	log.Info().
		Str("status_code", "204").
		Str("cart_id", c.Param("id")).
		Msg("Edit cart request processed successfully")

	c.Status(http.StatusNoContent)
}

// DiscardCartByID    godoc
// @Summary           Discard cart data
// @Description       delete cart data
// @Tags              carts
// @Accept            json
// @Produce           json
// @Param             Authorization  header  string  true  "Insert your admin JWT token"
// @Success           204
// @Failure           400  {object}  model.Error
// @Failure           500  {object}  model.Error
// @Router            /carts/:id [delete]
func (controller CartController) DiscardCartByID(c *gin.Context) {
	cartID := c.Param("id")

	_, err := controller.CartService.RemoveCartByID(&entity.Cart{
		ID: cartID,
	})
	if err != nil {
		log.Error().
			Err(err).
			Str("component", "controller.CartService.RemoveCartByID").
			Str("cart_id", cartID).
			Msg("Failed to fetch cart request")

		c.JSON(http.StatusInternalServerError, model.Error{Error: "discard_cart_failed", Message: "InternalServerError"})
		return
	}

	log.Info().
		Str("status_code", "204").
		Str("cart_id", cartID).
		Msg("Discard cart request processed successfully")

	c.Status(http.StatusNoContent)
}
