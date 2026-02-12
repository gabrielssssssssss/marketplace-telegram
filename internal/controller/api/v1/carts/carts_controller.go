package carts

import (
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/gabrielssssssssss/marketplace-telegram/helper"
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

// Create    	 godoc
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
func (controller CartController) Create(c *gin.Context) {
	var req model.Cart
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, model.Error{Error: "invalid_request", Message: "StatusBadRequest"})
		return
	}

	authorization := c.GetHeader("Authorization")

	userID, err := helper.GetUserID(authorization, os.Getenv("JWT_SECRET_KEY"))
	if err != nil {
		log.Error().
			Err(err).
			Str("component", "controller.UserService.GetUserID").
			Int64("user_id", userID).
			Msg("Failed to fetch user request")

		c.JSON(http.StatusUnauthorized, model.Error{Error: "invalid_token", Message: "Unauthorized"})
		return
	}

	cart, err := controller.CartService.Register(&entity.Cart{
		UserID:    userID,
		ProductID: req.ProductID,
	})
	if err != nil {
		log.Error().
			Err(err).
			Str("component", "controller.CartService.Register").
			Msg("Failed to insert cart request")

		c.JSON(http.StatusInternalServerError, model.Error{Error: "insert_cart_failed", Message: "InternalServerError"})
		return
	}

	log.Info().
		Str("status_code", "201").
		Msg("Insert cart request processed successfully")

	c.JSON(http.StatusCreated, model.CartResponse{Message: "success", Data: *cart})
}

// Carts    		godoc
// @Summary         Get carts data
// @Description     get carts data
// @Tags            carts
// @Accept          json
// @Produce         json
// @Param           Authorization  header  string  true  "Insert your admin JWT token"
// @Success         200  {object}  model.CartsResponse
// @Failure         500  {object}  model.Error
// @Router          /carts [get]
func (controller CartController) Carts(c *gin.Context) {
	authorization := c.GetHeader("Authorization")

	userID, err := helper.GetUserID(authorization, os.Getenv("JWT_SECRET_KEY"))
	if err != nil {
		log.Error().
			Err(err).
			Str("component", "controller.UserService.GetUserID").
			Int64("user_id", userID).
			Msg("Failed to fetch user request")

		c.JSON(http.StatusUnauthorized, model.Error{Error: "invalid_token", Message: "Unauthorized"})
		return
	}

	cart, err := controller.CartService.GetCarts(&entity.Cart{
		UserID: userID,
	})
	if err != nil {
		log.Error().
			Err(err).
			Str("component", "controller.CartService.GetCarts").
			Int64("user_id", userID).
			Msg("Failed to fetch carts request")

		c.JSON(http.StatusInternalServerError, model.Error{Error: "find_carts_failed", Message: "InternalServerError"})
		return
	}

	log.Info().
		Str("status_code", "200").
		Int64("user_id", userID).
		Msg("Fetch carts request processed successfully")

	c.JSON(http.StatusOK, model.CartsResponse{Message: "success", Data: *cart})
}

// Cart    			godoc
// @Summary         Get cart data
// @Description     get cart data
// @Tags            carts
// @Accept          json
// @Produce         json
// @Param           Authorization  header  string  true  "Insert your admin JWT token"
// @Success         200  {object}  model.CartResponse
// @Failure         500  {object}  model.Error
// @Router          /carts/:id [get]
func (controller CartController) Cart(c *gin.Context) {
	cartID := c.Param("id")

	cart, err := controller.CartService.GetCart(&entity.Cart{
		ID: cartID,
	})
	if err != nil {
		log.Error().
			Err(err).
			Str("component", "controller.CartService.GetCart").
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

// UserCarts  		   godoc
// @Summary            Get user carts data
// @Description        get all carts from userID data
// @Tags               carts
// @Accept             json
// @Produce            json
// @Param              Authorization  header  string  true  "Insert your admin JWT token"
// @Success            200  {object}  model.CartsResponse
// @Failure            500  {object}  model.Error
// @Router             /users/:id/carts [get]
func (controller CartController) UserCarts(c *gin.Context) {
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

	carts, err := controller.CartService.GetCarts(&entity.Cart{
		UserID: userID,
	})
	if err != nil {
		log.Error().
			Err(err).
			Str("component", "controller.CartService.GetCarts").
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

// Update     		godoc
// @Summary         Edit cart data
// @Description     put cart data
// @Tags            carts
// @Accept          json
// @Produce         json
// @Param           Authorization  header  string  true  "Insert your admin JWT token"
// @Success         200  {object}  model.CartResponse
// @Failure         400  {object}  model.Error
// @Failure         500  {object}  model.Error
// @Router          /carts/:id [put]
func (controller CartController) Update(c *gin.Context) {
	var req model.Cart
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, model.Error{Error: "invalid_request", Message: "StatusBadRequest"})
		return
	}

	cart, err := controller.CartService.Modify(&entity.Cart{
		ID:        c.Param("id"),
		UserID:    req.UserID,
		ProductID: req.ProductID,
		UpdatedAt: time.Now(),
	})
	if err != nil {
		log.Error().
			Err(err).
			Str("component", "controller.CartService.Modify").
			Str("cart_id", c.Param("id")).
			Msg("Failed to edit cart request")

		c.JSON(http.StatusInternalServerError, model.Error{Error: "edit_cart_failed", Message: "InternalServerError"})
		return
	}

	log.Info().
		Str("status_code", "200").
		Str("cart_id", c.Param("id")).
		Msg("Edit cart request processed successfully")

	c.JSON(http.StatusOK, model.CartResponse{Message: "success", Data: *cart})
}

// Delete    		  godoc
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
func (controller CartController) Delete(c *gin.Context) {
	cartID := c.Param("id")

	_, err := controller.CartService.Remove(&entity.Cart{
		ID: cartID,
	})
	if err != nil {
		log.Error().
			Err(err).
			Str("component", "controller.CartService.Remove").
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
