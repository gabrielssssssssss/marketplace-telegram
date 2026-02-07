package swagger

import (
	_ "github.com/gabrielssssssssss/marketplace-telegram/docs"
	"github.com/gin-gonic/gin"
	sf "github.com/swaggo/files"
	gs "github.com/swaggo/gin-swagger"
)

func Route(rg *gin.RouterGroup) {
	rg.GET("/swagger/*any", gs.WrapHandler(sf.Handler))
}
