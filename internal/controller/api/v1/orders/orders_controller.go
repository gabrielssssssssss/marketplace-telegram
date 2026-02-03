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

func (controller OrderController) InsertOrder(c *gin.Context) {
	var req model.Order
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid_request", "message": "StatusBadRequest"})
		return
	}

	order, err := controller.OrderService.RegisterOrder(&entity.Order{
		UserID:  req.UserID,
		Product: req.Product,
		Amount:  req.Amount,
	})
	if err != nil {
		log.Error().
			Err(err).
			Str("component", "controller.OrderService.RegisterOrder").
			Msg("Failed to insert order request")

		c.JSON(http.StatusInternalServerError, gin.H{"error": "insert_order_failed", "message": "InternalServerError"})
		return
	}

	log.Info().
		Str("status_code", "201").
		Msg("Insert order request processed successfully")

	c.JSON(http.StatusCreated, gin.H{"message": "success", "data": order})
}

func (controller OrderController) FetchOrderByID(c *gin.Context) {
	orderID := c.Param("id")

	order, err := controller.OrderService.GetOrderByID(&entity.Order{
		ID: orderID,
	})
	if err != nil {
		log.Error().
			Err(err).
			Str("component", "controller.OrderService.FindOrderByID").
			Str("order_id", orderID).
			Msg("Failed to fetch order request")

		c.JSON(http.StatusInternalServerError, gin.H{"error": "find_order_failed", "message": "InternalServerError"})
		return
	}

	log.Info().
		Str("status_code", "200").
		Str("cart_id", orderID).
		Msg("Fetch order request processed successfully")

	c.JSON(http.StatusOK, gin.H{"message": "success", "data": order})
}

func (controller OrderController) FetchOrdersByUserID(c *gin.Context) {
	params := c.Param("id")

	userID, err := strconv.ParseInt(params, 10, 64)
	if err != nil {
		log.Error().
			Err(err).
			Str("component", "strconv.ParseInt").
			Int64("user_id", userID).
			Msg("Failed to fetch order request")

		c.JSON(http.StatusInternalServerError, gin.H{"error": "find_order_failed", "message": "InternalServerError"})
		return
	}

	orders, err := controller.OrderService.GetOrdersByUserID(&entity.Order{
		UserID: userID,
	})
	if err != nil {
		log.Error().
			Err(err).
			Str("component", "controller.OrderService.FindOrdersByUserID").
			Int64("user_id", userID).
			Msg("Failed to fetch order request")

		c.JSON(http.StatusInternalServerError, gin.H{"error": "find_order_failed", "message": "InternalServerError"})
		return
	}

	log.Info().
		Str("status_code", "200").
		Int64("order_id", userID).
		Msg("Fetch order request processed successfully")

	c.JSON(http.StatusOK, gin.H{"message": "success", "data": orders})
}

func (controller OrderController) DiscardOrderByID(c *gin.Context) {
	orderID := c.Param("id")

	_, err := controller.OrderService.RemoveOrderByID(&entity.Order{
		ID: orderID,
	})
	if err != nil {
		log.Error().
			Err(err).
			Str("component", "controller.OrderService.RemoveOrderByID").
			Str("order_id", orderID).
			Msg("Failed to fetch order request")

		c.JSON(http.StatusInternalServerError, gin.H{"error": "discard_order_failed", "message": "InternalServerError"})
		return
	}

	log.Info().
		Str("status_code", "204").
		Str("cart_id", orderID).
		Msg("Discard order request processed successfully")

	c.Status(http.StatusNoContent)
}
