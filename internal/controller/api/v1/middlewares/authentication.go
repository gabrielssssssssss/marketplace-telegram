package middlewares

import (
	"net/http"
	"os"

	"github.com/gabrielssssssssss/marketplace-telegram/helper"
	"github.com/gabrielssssssssss/marketplace-telegram/internal/model"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

func Authorization(c *gin.Context) {
	authorization := c.GetHeader("Authorization")
	if authorization == "" {
		c.AbortWithStatusJSON(http.StatusUnauthorized, model.Error{Error: "missing_token", Message: "Unauthorized"})
		return
	}

	token, err := helper.VerifyJwtToken(authorization, os.Getenv("JWT_SECRET_KEY"))
	if err != nil || !token.Valid {
		log.Error().
			Err(err).
			Str("component", "helper.VerifyJwtToken").
			Str("session", authorization).
			Msg("Failed to process verify session token")

		c.AbortWithStatusJSON(http.StatusUnauthorized, model.Error{Error: "invalid_token", Message: "Unauthorized"})
		return
	}

	c.Next()
}
