package main

import (
	"fmt"
	"server/core"
	"server/global"
)

func main() {
	// 初始化Viper
	global.GViper = core.Viper("./config.yaml")
	if err := core.InitLogger(&global.GConfig); err != nil {
		fmt.Printf("init logger failed, err:%v\n", err)
		return
	}
	router := core.SetupRouter()
	if err := router.Run(":" + global.GConfig.System.Port); err != nil {
		fmt.Printf("startup service failed, err:%v\n\n", err)
	}
}
