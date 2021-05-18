package router

import (
	"github.com/gin-gonic/gin"
	"server/api"
)

func InitUserRouter(Router *gin.RouterGroup) {
	UserRouter := Router.Group("user")
	{
		UserRouter.POST("/AccountLogin", api.AccountLogin)
		UserRouter.POST("/register", api.Register)
	}
}
