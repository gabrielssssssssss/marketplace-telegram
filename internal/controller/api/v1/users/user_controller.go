package users

import (
	"net/http"

	"github.com/gabrielssssssssss/marketplace-telegram/helper"
	"github.com/gabrielssssssssss/marketplace-telegram/internal/service"
	"github.com/gin-gonic/gin"
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
		c.JSON(http.StatusUnauthorized, gin.H{
			"error":   "invalid_token",
			"message": err.Error(),
		})
		return
	}

	user, err := controller.AccountService.FindUserByID(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "find_user_failed",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "user found with successfully!",
		"data":    user,
	})
}
