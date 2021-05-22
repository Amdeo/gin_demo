package api

import (
	"github.com/gin-gonic/gin"
)

const (
	PhoneLogin   = "0"
	AccountLogin = "1"
)

func Login(c *gin.Context) {
	// 获取url参数
	loginMethod := c.DefaultQuery("method", "")

	if loginMethod == PhoneLogin {
		phoneLogin(c)
	} else if loginMethod == AccountLogin {
		accountLogin(c)
	}
}

func phoneLogin(c *gin.Context) {
	//var user = new(model.User)
	//phone := c.DefaultQuery("phone", "")
	//auth_code := c.DefaultQuery("auth_code", "")
	//g.DB.Where("Phone = ?", phone).First(user)
	//zap.L().Debug("%v",user)
	//fmt.Printf("%+v", user)
	//common.OkWithData(,c)

}

func accountLogin(c *gin.Context) {

}
