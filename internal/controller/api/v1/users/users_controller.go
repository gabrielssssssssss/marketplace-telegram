package users

import (
	"net/http"
	"os"
	"time"

	"github.com/gabrielssssssssss/marketplace-telegram/helper"
	"github.com/gabrielssssssssss/marketplace-telegram/internal/entity"
	"github.com/gabrielssssssssss/marketplace-telegram/internal/model"
	"github.com/gabrielssssssssss/marketplace-telegram/internal/service"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

type UserController struct {
	UserService service.UserService
}

func NewUserController(UserService *service.UserService) UserController {
	return UserController{UserService: *UserService}
}

func (controller UserController) FetchUserByID(c *gin.Context) {
	authorization := c.GetHeader("Authorization")

	userID, err := helper.GetUserID(authorization, os.Getenv("JWT_SECRET_KEY"))
	if err != nil {
		log.Error().
			Err(err).
			Str("component", "controller.UserService.FindUserByID").
			Int64("user_id", userID).
			Msg("Failed to fetch user request")

		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid_token", "message": "Unauthorized"})
		return
	}

	user, err := controller.UserService.GetUserByID(&entity.User{UserId: userID})
	if err != nil {
		log.Error().
			Err(err).
			Str("component", "controller.UserService.FindUserByID").
			Int64("user_id", userID).
			Msg("Failed to fetch user request")

		c.JSON(http.StatusInternalServerError, gin.H{"error": "find_user_failed", "message": "InternalServerError"})
		return
	}

	log.Info().
		Str("status_code", "200").
		Int64("user_id", userID).
		Msg("Fetch user request processed successfully")

	c.JSON(http.StatusOK, gin.H{"message": "success", "data": user})
}

func (controller UserController) DiscardUserByID(c *gin.Context) {
	var req model.User
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid_request", "message": "StatusBadRequest"})
		return
	}

	_, err := controller.UserService.RemoveUserByID(&entity.User{UserId: req.UserId})
	if err != nil {
		log.Error().
			Err(err).
			Str("component", "controller.UserService.RemoveUserByID").
			Int64("user_id", req.UserId).
			Msg("Failed to discard user request")

		c.JSON(http.StatusInternalServerError, gin.H{"error": "discard_user_failed", "message": "InternalServerError"})
		return
	}

	log.Info().
		Str("status_code", "204").
		Int64("user_id", req.UserId).
		Msg("Discard user request processed successfully")

	c.Status(http.StatusNoContent)
}

func (controller UserController) EditUserByID(c *gin.Context) {
	var req model.User
	if err := c.ShouldBindJSON(&req); err != nil || req.UserId == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid_request", "message": "StatusBadRequest"})
		return
	}

	_, err := controller.UserService.ModifyUserByID(&entity.User{
		UserId:      req.UserId,
		Firstname:   req.Firstname,
		Lastname:    req.Lastname,
		Username:    req.Username,
		Balance:     req.Balance,
		RecoveryKey: req.RecoveryKey,
		UpdatedAt:   time.Now(),
	})
	if err != nil {
		log.Error().
			Err(err).
			Str("component", "controller.UserService.ModifyUserByID").
			Int64("user_id", req.UserId).
			Msg("Failed to discard user request")

		c.JSON(http.StatusInternalServerError, gin.H{"error": "edit_user_failed", "message": "InternalServerError"})
		return
	}

	log.Info().
		Str("status_code", "204").
		Int64("user_id", req.UserId).
		Msg("Edit user request processed successfully")

	c.Status(http.StatusNoContent)
}
