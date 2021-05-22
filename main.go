package main

import (
	"fmt"
	"server/core"
	g "server/global"
)

func main() {
	// 初始化Viper
	g.Viper = core.Viper("./config.yaml")

	if err := core.InitLogger(&g.Config); err != nil {
		panic(fmt.Sprintf("init logger failed, err:%v\n", err))
		return
	}
	core.RunServer()
}
