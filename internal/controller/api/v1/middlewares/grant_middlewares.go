package middlewares

import (
	"net/http"

	"github.com/gabrielssssssssss/marketplace-telegram/helper"
	"github.com/gabrielssssssssss/marketplace-telegram/internal/service"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

func Grant(accountService *service.AccountService) gin.HandlerFunc {
	return func(c *gin.Context) {
		authorization := c.GetHeader("Authorization")
		if authorization == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "missing_token", "message": "Unauthorized"})
			return
		}

		userID, err := helper.GetJwtValue(authorization, "user_id")
		if err != nil {
			log.Error().
				Err(err).
				Str("component", "helper.GetJwtValue").
				Int64("user_id", userID).
				Msg("Failed to give grant access for user request")

			c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid_token", "message": "Unauthorized"})
			return
		}

		user, err := accountService.FindUserByID(userID)
		if err != nil {
			log.Error().
				Err(err).
				Str("component", "accountService.FindUserByID").
				Int64("user_id", userID).
				Msg("Failed found user_id")

			c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid_token", "message": "Unauthorized"})
			return
		}

		if user.Role == "admin" {
			c.Next()
		}
	}
}
