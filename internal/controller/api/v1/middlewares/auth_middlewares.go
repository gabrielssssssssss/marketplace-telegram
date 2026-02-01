package middlewares

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gabrielssssssssss/marketplace-telegram/helper"
	"github.com/gin-gonic/gin"
)

func Authentification(c *gin.Context) {
	token := c.GetHeader("Authorization")
	if token == "" {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"error":   "missing_token",
			"message": "Authorization header is absent or empty",
		})
		return
	}

	jwt, err := helper.VerifyJwtToken(token, os.Getenv("JWT_SECRET_KEY"))
	if err != nil || !jwt.Valid {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"error":   "invalid_token",
			"message": fmt.Sprintf("JWT validation failed: %v", err),
		})
		return
	}

	c.Next()
}
