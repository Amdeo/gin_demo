package router

import (
	"github.com/gin-gonic/gin"
	"server/apps/Rc/api"
)

func Router(Router *gin.RouterGroup) {
	InitUserRouter(Router)
}

func InitUserRouter(Router *gin.RouterGroup) {
	UserRouter := Router.Group("user")
	{
		UserRouter.GET("/login", api.Login)
	}
}
