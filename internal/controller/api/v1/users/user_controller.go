package users

import (
	"github.com/gabrielssssssssss/marketplace-telegram/internal/service"
	"github.com/gin-gonic/gin"
)

type UserController struct {
	AccountService *service.AccountService
}

func NewUserController(AccountService *service.AccountService) UserController {
	return UserController{AccountService: AccountService}
}

func (controller UserController) FetchUserByID(c *gin.Context) {

}
