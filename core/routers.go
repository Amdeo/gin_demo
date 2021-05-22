package core

import (
	RcRouter "server/apps/Rc/router"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	Router := gin.Default()
	// 增加
	Router.Use(GinLogger(), GinRecovery(true))

	RcGroup := Router.Group("test")
	{
		RcRouter.Router(RcGroup)
	}
	return Router
}
