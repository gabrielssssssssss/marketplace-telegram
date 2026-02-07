package users

import (
	"github.com/gabrielssssssssss/marketplace-telegram/internal/controller/api/v1/middlewares"
	"github.com/gin-gonic/gin"
)

func (controller *UserController) Route(rg *gin.RouterGroup) {
	rg.GET("/users",
		middlewares.Authorization,
		controller.FetchUserByID,
	)
	rg.DELETE("/users",
		middlewares.Authorization,
		middlewares.Grant(),
		controller.DiscardUserByID,
	)
	rg.PUT("/users",
		middlewares.Authorization,
		middlewares.Grant(),
		controller.EditUserByID,
	)
}
