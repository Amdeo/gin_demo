package global

import (
	"github.com/go-redis/redis"
	"github.com/spf13/viper"
	"gorm.io/gorm"
	"server/config"
)

var (
	GConfig config.ServerConfig
	GViper  *viper.Viper
	GDb     *gorm.DB
	GRedis  *redis.Client
)
