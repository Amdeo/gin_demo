package core

import (
	"github.com/gin-gonic/gin"
	"server/router"
)

func SetupRouter() *gin.Engine {
	Router := gin.Default()
	// 增加
	Router.Use(GinLogger(), GinRecovery(true))
	PrivateGroup := Router.Group("")
	{
		router.InitUserRouter(PrivateGroup)
	}
	return Router
}
