package global

import (
	"github.com/go-redis/redis"
	"github.com/spf13/viper"
	"gorm.io/gorm"
	"server/config"
)

var (
	Config config.ServerConfig
	Viper  *viper.Viper
	DB     *gorm.DB
	Redis  *redis.Client
)

//type globalVar struct {
//	conf config.ServerConfig
//	viper *viper.Viper
//	db *gorm.DB
//	cache *redis.Client
//}
//
//var G = new(globalVar)
