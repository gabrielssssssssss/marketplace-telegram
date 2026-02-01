package middlewares

import (
	"net/http"
	"os"

	"github.com/gabrielssssssssss/marketplace-telegram/helper"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

func Authentification(c *gin.Context) {
	authorization := c.GetHeader("Authorization")
	if authorization == "" {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "missing_token", "message": "Unauthorized"})
		return
	}

	token, err := helper.VerifyJwtToken(authorization, os.Getenv("JWT_SECRET_KEY"))
	if err != nil || !token.Valid {
		log.Error().
			Err(err).
			Str("component", "helper.VerifyJwtToken").
			Str("session", authorization).
			Msg("Failed to process verify session token")

		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "invalid_token", "message": "Unauthorized"})
		return
	}

	c.Next()
}
