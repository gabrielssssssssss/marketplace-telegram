package users

import (
	"github.com/gabrielssssssssss/marketplace-telegram/internal/controller/api/v1/middlewares"
	"github.com/gin-gonic/gin"
)

func (controller *UserController) Route(rg *gin.RouterGroup) {
	rg.GET("/users", controller.FetchUserByID)
	rg.DELETE("/users", middlewares.Grant(), controller.DiscardUserByID)
	rg.PUT("/users", middlewares.Grant(), controller.EditUserByID)
}
