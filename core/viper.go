package core

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	g "server/global"
)

func Viper(path ...string) *viper.Viper {
	var config string
	if len(path) == 0 {
		fmt.Printf("config path error")
	}
	config = path[0]

	v := viper.New()
	// 设置配置文件路径
	v.SetConfigFile(config)
	// 配置文件类型
	v.SetConfigType("yaml")
	// 读取配置文件
	err := v.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
	// 监控配置文件变化
	v.WatchConfig()
	v.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("config file changed:", e.Name)
		if err := v.Unmarshal(&g.Config); err != nil {
			fmt.Println(err)
		}
	})
	if err := v.Unmarshal(&g.Config); err != nil {
		fmt.Println(err)
	}
	return v
}
