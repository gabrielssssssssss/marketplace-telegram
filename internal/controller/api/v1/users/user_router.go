package users

import "github.com/gin-gonic/gin"

func (controller *UserController) Route(rg *gin.RouterGroup) {
	rg.GET("/users", controller.FetchUserByID)
	rg.DELETE("/users", controller.DiscardUserByID)
	rg.PUT("/users", controller.EditUserByID)
}
