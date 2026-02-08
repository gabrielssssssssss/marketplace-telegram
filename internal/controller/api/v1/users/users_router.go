package users

import (
	"github.com/gabrielssssssssss/marketplace-telegram/internal/controller/api/v1/middlewares"
	"github.com/gin-gonic/gin"
)

func (controller *UserController) Route(rg *gin.RouterGroup) {
	rg.GET("/me",
		middlewares.Authorization,
		controller.Me,
	)
	users := rg.Group("/users", middlewares.Authorization, middlewares.Grant())
	{
		users.GET("/:id", controller.FetchUserByID)
		users.PUT("/:id", controller.EditUserByID)
		users.DELETE("/:id", controller.DiscardUserByID)
	}
}
