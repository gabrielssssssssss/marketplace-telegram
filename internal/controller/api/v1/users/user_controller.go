package users

import (
	"net/http"

	"github.com/gabrielssssssssss/marketplace-telegram/helper"
	"github.com/gabrielssssssssss/marketplace-telegram/internal/service"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

type UserController struct {
	AccountService service.AccountService
}

func NewUserController(AccountService *service.AccountService) UserController {
	return UserController{AccountService: *AccountService}
}

func (controller UserController) FetchUserByID(c *gin.Context) {
	authorization := c.GetHeader("Authorization")

	userID, err := helper.GetJwtValue(authorization, "user_id")
	if err != nil {
		log.Error().
			Err(err).
			Str("component", "controller.AccountService.FindUserByID").
			Int64("user_id", userID).
			Msg("Failed to fetch user request")

		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid_token", "message": "Unauthorized"})
		return
	}

	user, err := controller.AccountService.FindUserByID(userID)
	if err != nil {
		log.Error().
			Err(err).
			Str("component", "controller.AccountService.FindUserByID").
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
