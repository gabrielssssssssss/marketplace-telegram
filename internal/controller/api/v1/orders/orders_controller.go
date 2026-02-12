package orders

import (
	"net/http"
	"strconv"

	"github.com/gabrielssssssssss/marketplace-telegram/internal/entity"
	"github.com/gabrielssssssssss/marketplace-telegram/internal/model"
	"github.com/gabrielssssssssss/marketplace-telegram/internal/service"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

type OrderController struct {
	OrderService service.OrderService
}

func NewOrderController(OrderService *service.OrderService) OrderController {
	return OrderController{OrderService: *OrderService}
}

// Register      godoc
// @Summary      Insert order data
// @Description  post order data
// @Tags         orders
// @Accept       json
// @Produce      json
// @Param        Authorization  header  string  true  "Insert your admin JWT token"
// @Success      201  {object}  model.OrderResponse
// @Failure      400  {object}  model.Error
// @Failure      500  {object}  model.Error
// @Router         /orders/ [post]
func (controller OrderController) Register(c *gin.Context) {
	var req model.Order
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid_request", "message": "StatusBadRequest"})
		return
	}

	order, err := controller.OrderService.Register(&entity.Order{
		UserID:  req.UserID,
		Product: req.Product,
		Amount:  req.Amount,
	})
	if err != nil {
		log.Error().
			Err(err).
			Str("component", "controller.OrderService.Register").
			Msg("Failed to insert order request")

		c.JSON(http.StatusInternalServerError, gin.H{"error": "insert_order_failed", "message": "InternalServerError"})
		return
	}

	log.Info().
		Str("status_code", "201").
		Msg("Insert order request processed successfully")

	c.JSON(http.StatusCreated, model.OrderResponse{Message: "success", Data: *order})
}

// GetOrder   		godoc
// @Summary         Get order data
// @Description     get order data
// @Tags            orders
// @Accept          json
// @Produce         json
// @Param           Authorization  header  string  true  "Insert your admin JWT token"
// @Success         200  {object}  model.OrderResponse
// @Failure         500  {object}  model.Error
// @Router          /orders/:id [get]
func (controller OrderController) GetOrder(c *gin.Context) {
	orderID := c.Param("id")

	order, err := controller.OrderService.GetOrder(&entity.Order{
		ID: orderID,
	})
	if err != nil {
		log.Error().
			Err(err).
			Str("component", "controller.OrderService.GetOrder").
			Str("order_id", orderID).
			Msg("Failed to fetch order request")

		c.JSON(http.StatusInternalServerError, model.Error{Error: "find_order_failed", Message: "InternalServerError"})
		return
	}

	log.Info().
		Str("status_code", "200").
		Str("cart_id", orderID).
		Msg("Fetch order request processed successfully")

	c.JSON(http.StatusOK, model.OrderResponse{Message: "success", Data: *order})
}

// GetUserOrders 	   godoc
// @Summary            Get user orders data
// @Description        get all orders from userID data
// @Tags               orders
// @Accept             json
// @Produce            json
// @Param              Authorization  header  string  true  "Insert your admin JWT token"
// @Success            200  {object}  model.OrdersResponse
// @Failure            500  {object}  model.Error
// @Router             /users/:id/orders [get]
func (controller OrderController) GetUserOrders(c *gin.Context) {
	params := c.Param("id")

	userID, err := strconv.ParseInt(params, 10, 64)
	if err != nil {
		log.Error().
			Err(err).
			Str("component", "strconv.ParseInt").
			Int64("user_id", userID).
			Msg("Failed to fetch order request")

		c.JSON(http.StatusInternalServerError, model.Error{Error: "find_order_failed", Message: "InternalServerError"})
		return
	}

	orders, err := controller.OrderService.GetUserOrders(&entity.Order{
		UserID: userID,
	})
	if err != nil {
		log.Error().
			Err(err).
			Str("component", "controller.OrderService.GetUserOrders").
			Int64("user_id", userID).
			Msg("Failed to fetch order request")

		c.JSON(http.StatusInternalServerError, model.Error{Error: "find_order_failed", Message: "InternalServerError"})
		return
	}

	log.Info().
		Str("status_code", "200").
		Int64("order_id", userID).
		Msg("Fetch order request processed successfully")

	c.JSON(http.StatusOK, model.OrdersResponse{Message: "success", Data: *orders})
}

// Remove   		  godoc
// @Summary           Discard order data
// @Description       delete order data
// @Tags              orders
// @Accept            json
// @Produce           json
// @Param             Authorization  header  string  true  "Insert your admin JWT token"
// @Success           204
// @Failure           400  {object}  model.Error
// @Failure           500  {object}  model.Error
// @Router            /orders/:id [delete]
func (controller OrderController) Remove(c *gin.Context) {
	orderID := c.Param("id")

	_, err := controller.OrderService.Remove(&entity.Order{
		ID: orderID,
	})
	if err != nil {
		log.Error().
			Err(err).
			Str("component", "controller.OrderService.Remove").
			Str("order_id", orderID).
			Msg("Failed to fetch order request")

		c.JSON(http.StatusInternalServerError, model.Error{Error: "discard_order_failed", Message: "InternalServerError"})
		return
	}

	log.Info().
		Str("status_code", "204").
		Str("cart_id", orderID).
		Msg("Discard order request processed successfully")

	c.Status(http.StatusNoContent)
}
