package middlewares

import (
	"net/http"
	"os"

	"github.com/gabrielssssssssss/marketplace-telegram/helper"
	"github.com/gabrielssssssssss/marketplace-telegram/internal/model"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

func Grant() gin.HandlerFunc {
	return func(c *gin.Context) {
		authorization := c.GetHeader("Authorization")
		if authorization == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, model.Error{Error: "missing_token", Message: "Unauthorized"})
			return
		}

		userID, err := helper.GetUserID(authorization, os.Getenv("JWT_SECRET_KEY"))
		if err != nil {
			log.Error().
				Err(err).
				Str("component", "helper.GetUserID").
				Int64("user_id", userID).
				Msg("Failed to give grant access for user request")

			c.JSON(http.StatusUnauthorized, model.Error{Error: "invalid_token", Message: "Unauthorized"})
			return
		}

		userRole, err := helper.GetRole(authorization, os.Getenv("JWT_SECRET_KEY"))
		if err != nil {
			log.Error().
				Err(err).
				Str("component", "helper.GetRole").
				Int64("user_id", userID).
				Msg("Failed to give grant access for user request")

			c.JSON(http.StatusUnauthorized, model.Error{Error: "invalid_token", Message: "Unauthorized"})
			return
		}

		if userRole == "admin" {
			c.Next()
		}
	}
}
